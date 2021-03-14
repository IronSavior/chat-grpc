package chat

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IronSavior/chat-grpc/chatpb/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbi "google.golang.org/protobuf/runtime/protoiface"
)

// Chat channel buffer length
const msgQueueSize = 10

const shutdownTimeout = 10 * time.Second

// Service owns service-wide state and connects new client streams to specific chats
type Service struct {
	// Running chats
	reg registry

	// Incoming command queue
	commands UserCommands

	// stop is closed when the service is shutting down or stopped
	stop  chan struct{}
	stop1 sync.Once

	// done is closed when the service and its chats have completely shut down
	done chan struct{}

	// 0 = Run() not yet called; 1 = Run() was started
	started int64
}

// New initializes a new service object. qSize is the max depth of the service-level
// command queue (as opposed to chat command queues)
func New(qSize int) *Service {
	return &Service{
		commands: make(chan UserCommand, qSize),
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}
}

// Run drives the main service goroutine. Must be called only once.
func (s *Service) Run(ctx context.Context) error {
	if ok := atomic.CompareAndSwapInt64(&s.started, 0, 1); !ok {
		return status.Errorf(codes.Internal, "BUG: chat service already started once")
	}
	defer close(s.done)

	err := s.cmdLoop(ctx)

	// Block and drain command queue
	cmds := s.commands
	s.commands = nil

	drainErr := status.Error(codes.Aborted, "chat service is not available on this node")
	for _, uc := range cmds.RejectAll(drainErr) {
		uc.User.CloseEvents(drainErr)
	}

	s.shutdownChats(shutdownTimeout)
	return err
}

func (s *Service) shutdownChats(timeout time.Duration) error {
	wg := sync.WaitGroup{}
	chats := s.reg.Chats()
	wg.Add(len(chats))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, chat := range chats {
		go func(chat *Chat) {
			defer wg.Done()

			// TODO: detect timeout during shutdown and do something about them
			chat.Shutdown(ctx)
		}(chat)
	}

	wg.Wait()
	return nil
}

func (s *Service) cmdLoop(ctx context.Context) error {
	defer s.Stop()

	for {
		select {
		case uc, ok := <-s.commands:
			if !ok {
				return nil
			}

			uc.User.CloseEvents(s.handleCommand(ctx, uc))

		case <-ctx.Done():
			return status.FromContextError(ctx.Err()).Err()

		case <-s.stop:
			return nil
		}
	}
}

func (s *Service) handleCommand(ctx context.Context, uc UserCommand) error {
	switch cmd := uc.Command.(type) {
	case *commands.Create:
		chat := NewChat(cmd.ChatName, msgQueueSize)
		go chat.Run(ctx)

		if err := s.reg.Add(cmd.ChatName, chat); err != nil {
			chat.Stop()
			return err
		}

		return nil

	case *commands.Destroy:
		chat, _ := s.reg.Get(cmd.ChatName)
		if err := s.reg.Remove(cmd.ChatName); err != nil {
			return err
		}

		chat.Stop()
		return nil

	default:
		return status.Errorf(codes.Internal, "BUG: unexpected command type: %T", uc.Command)
	}
}

// invokeCmdSync sends a command into the service and waits for its response
func (s *Service) invokeCmdSync(ctx context.Context, cmd pbi.MessageV1, userName string) (pbi.MessageV1, error) {
	assertIsCommand(cmd)
	stream := ResponseStream{events: make(chan pbi.MessageV1)}

	uc := UserCommand{
		User: User{
			name:   userName,
			Events: Events{&stream},
		},
		Command: cmd,
	}

	select {
	case <-s.stop:
		return nil, status.Error(codes.Unavailable, "chat service is not available on this node")

	case <-ctx.Done():
		return nil, status.FromContextError(ctx.Err()).Err()

	case s.commands <- uc:
	}

	var event pbi.MessageV1
	for {
		select {
		case <-s.stop:
			return nil, status.Error(codes.Unavailable, "chat service is not available on this node")

		case <-ctx.Done():
			return nil, status.FromContextError(ctx.Err()).Err()

		case evt, ok := <-stream.events:
			if !ok {
				// CloseEvents was called
				return event, stream.Err()
			}

			if event != nil {
				// expectation is the service sets stream error and closes the stream after sending the first event
				return nil, status.Errorf(codes.Internal, "BUG: service returned more than one event (zero or one expected)")
			}

			event = evt
		}
	}
}

// CreateChat creates a new chat. Returns events.Created with nil error or events.Rejected with error
func (s *Service) CreateChat(ctx context.Context, userName, chatName string, qSize int) (pbi.MessageV1, error) {
	switch {
	case chatName == "":
		return nil, status.Errorf(codes.InvalidArgument, "invalid chat name %q", chatName)
	}
	// TODO: other name validations

	return s.invokeCmdSync(ctx, &commands.Create{ChatName: chatName}, userName)
}

// DestroyChat destroys an existing chat. Returns events.Destoyed or events.Rejected
func (s *Service) DestroyChat(ctx context.Context, userName, chatName string) (pbi.MessageV1, error) {
	return s.invokeCmdSync(ctx, &commands.Destroy{ChatName: chatName}, userName)
}

// Stream connects a user client stream to a chat
func (s *Service) Stream(ctx context.Context, userName, chatName string, stream Stream) error {
	switch {
	case userName == "":
		return status.Error(codes.InvalidArgument, "invalid user name")
		// TODO: what is a valid name?
	}

	// Wait for user to request join
	var cmd pbi.MessageV1
	select {
	case <-s.stop:
		return status.Error(codes.Unavailable, "chat service is not available on this node")

	case <-ctx.Done():
		return status.FromContextError(ctx.Err()).Err()

	case c, ok := <-stream.Commands():
		if !ok {
			return stream.Err()
		}
		cmd = c
	}

	join, ok := cmd.(*commands.Join)
	if !ok {
		return status.Error(codes.InvalidArgument, "expected a join command")
	}

	chat, ok := s.reg.Get(chatName)
	if !ok {
		return status.Errorf(codes.NotFound, "chat %q doesn't exist", chatName)
	}

	return chat.Stream(ctx, userName, stream, join)
}

func (s *Service) Done() <-chan struct{} {
	return s.done
}

// Stop signals the chat to shut down (non blocking)
func (s *Service) Stop() {
	s.stop1.Do(func() { close(s.stop) })
}

// Shutdown signals the chat to stop and blocks until it stops or the context is cancelled
func (s *Service) Shutdown(ctx context.Context) error {
	s.Stop()

	select {
	case <-ctx.Done():
		return status.FromContextError(ctx.Err()).Err()
	case <-s.done:
		return nil
	}
}
