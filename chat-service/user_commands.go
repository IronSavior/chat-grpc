package chat

import (
	"time"
)

// UserCommands is a channel of UserCommand with some added behavior
type UserCommands chan UserCommand

// RejectAll drains the channel and rejects any commands found. Returns slice of drained commands.
func (ucs UserCommands) RejectAll(err error) []UserCommand {
	var ucmds []UserCommand

	for {
		select {
		case uc, ok := <-ucs:
			if !ok {
				return ucmds
			}

			ucmds = append(ucmds, uc)
			uc.TryReject(err) //  Don't care about un-received messages here

		default:
			return ucmds
		}
	}
}

// SendTimeout sends the event on the stream unless it would block for longer than the given timeout
// Returns true when event is sent
func (ucs UserCommands) SendTimeout(uc UserCommand, timeout time.Duration) bool {
	select {
	case ucs <- uc:
		return true
	case <-time.After(timeout):
		return false
	}
}
