package chat

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IronSavior/chat-grpc/chatpb/commands"
	"github.com/IronSavior/chat-grpc/chatpb/events"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbi "google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Chat is a broadcast domain among one group of participants
type Chat struct {
	// Name of this chat
	name string

	// Joined participants
	roster roster

	// Users whose event queues were full when the chat tried to send
	unresponsive []User

	// incoming commands
	commands UserCommands

	// stop is closed when the chat is shutting down or stopped
	stop  chan struct{}
	stop1 sync.Once

	// done is closed when the chat has completely shut down
	done chan struct{}

	// A Chat is only allowed to start once
	// 0 = Run() not yet called; 1 = Run() was started
	started int32
}

func NewChat(name string, qSize int) *Chat {
	return &Chat{
		name:     name,
		commands: make(chan UserCommand, qSize),
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}
}

// Run drives the chat
func (c *Chat) Run(ctx context.Context) error {
	if ok := atomic.CompareAndSwapInt32(&c.started, 0, 1); !ok {
		return status.Errorf(codes.Internal, "chat %q already started", c.name)
	}
	defer close(c.done)

	if err := c.notifyAll(&events.Created{}); err != nil {
		c.Stop()
		return err
	}

	err := c.cmdLoop(ctx)

	userErr := status.Error(codes.Aborted, "chat has ended")
	c.drainCommands(userErr)

	for _, u := range c.roster.Users() {
		c.notifyAll(unjoinedEvt(u, userErr))
		c.roster.Remove(u)
		u.CloseEvents(userErr)
	}
	c.purgeUnresponsive()

	return err
}

func (c *Chat) drainCommands(err error) {
	for {
		select {
		case uc, ok := <-c.commands:
			if !ok {
				return
			}

			if !c.roster.Has(uc.User) {
				continue
			}

			uc.TryReject(err) //  Don't care about un-received messages here

		default:
			return
		}
	}
}

func (c *Chat) cmdLoop(ctx context.Context) error {
	defer c.Stop()

	for {
		select {
		case uc, ok := <-c.commands:
			if !ok {
				return nil
			}

			if err := c.handleCommand(uc.User, uc.Command); err != nil {
				// TODO: handle error by something other than killing the chat
				return err
			}

		case <-ctx.Done():
			return ctx.Err()

		case <-c.stop:
			return nil
		}

		c.purgeUnresponsive()
	}
}

func (c *Chat) handleCommand(u User, cmd pbi.MessageV1) error {
	switch cmd := cmd.(type) {
	case *commands.Say:
		return c.notifyAll(&events.Said{
			Time:     timestamppb.Now(),
			UserName: u.Name(),
			Message:  cmd.Message,
		})

	case *commands.Join:
		if err := c.roster.Add(u); err != nil {
			return err
		}

		return c.notifyAll(&events.Joined{
			Time:     timestamppb.Now(),
			UserName: u.Name(),
		})

	case *commands.Unjoin:
		userErr := status.Error(codes.Code(cmd.Code), cmd.Message)
		c.notifyAll(unjoinedEvt(u, userErr))
		c.roster.Remove(u)
		u.CloseEvents(userErr)

		return nil

	default:
		return status.Errorf(codes.Internal, "BUG: unknown chat command type: %T", cmd)
	}
}

func (c *Chat) notifyAll(event pbi.MessageV1) error {
	if !IsEvent(event) {
		return status.Errorf(codes.Internal, "BUG: unknown chat event type: %T", event)
	}

	for _, u := range c.roster.Users() {
		if !u.TrySend(event) {
			c.roster.Remove(u)
			c.unresponsive = append(c.unresponsive, u)
		}
	}

	return nil
}

func (c *Chat) purgeUnresponsive() {
	for _, u := range c.unresponsive {
		err := status.Error(codes.DataLoss, "user unresponsive")
		u.CloseEvents(err)
		c.notifyAll(unjoinedEvt(u, err))
	}
	c.unresponsive = nil
}

// Done returns a signal channel that will close when the chat shuts down
func (c *Chat) Done() <-chan struct{} {
	return c.done
}

// Stop signals the chat to shut down (non blocking)
func (c *Chat) Stop() {
	c.stop1.Do(func() { close(c.stop) })
}

// Stopped returns a signal channel that is closed once the chat starts to shut down
func (c *Chat) Stopped() <-chan struct{} {
	return c.stop
}

// Shutdown signals the chat to stop and blocks until it stops or the context is cancelled
func (c *Chat) Shutdown(ctx context.Context) error {
	c.Stop()

	select {
	case <-ctx.Done():
		return status.FromContextError(ctx.Err()).Err()
	case <-c.done:
		return nil
	}
}

// sendUC writes a command to the Chat's command queue
func (c *Chat) sendUC(ctx context.Context, uc UserCommand) error {
	select {
	case <-c.stop:
		return status.Errorf(codes.Unavailable, "chat %q is stopped", c.name)

	case <-ctx.Done():
		return status.FromContextError(ctx.Err()).Err()

	case c.commands <- uc:
		return nil
	}
}

// Stream joins a user stream to this chat and pumps in messages from the user's command channel
func (c *Chat) Stream(ctx context.Context, userName string, stream Stream, join *commands.Join) error {
	u := User{
		name:          userName,
		Events:        Events{stream},
		CommandStream: stream,
	}

	// In error cases, try to send an unjoin command in last-ditch effort
	unjoin := func(err error) error {
		maxWait := 5 * time.Second
		uc := UserCommand{User: u, Command: unjoinCmd(err)}
		c.commands.SendTimeout(uc, maxWait)
		return err
	}

	// Send join cmd to chat
	if err := c.sendUC(ctx, UserCommand{User: u, Command: join}); err != nil {
		return err
	}

	// command loop
	for {
		select {
		case <-c.stop: // Chat shutting down, no more commands accepted
			return nil // chat will unjoin

		case <-ctx.Done(): // User stream cancelled
			return unjoin(status.FromContextError(ctx.Err()).Err())

		case cmd, ok := <-u.Commands():
			if !ok { // Command stream was closed, return stream error (if any)
				return unjoin(u.Err())
			}

			// Got a command, send it in
			if err := c.sendUC(ctx, UserCommand{User: u, Command: cmd}); err != nil {
				return unjoin(err)
			}
		}
	}
}
