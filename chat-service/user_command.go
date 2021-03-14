package chat

import (
	pbi "google.golang.org/protobuf/runtime/protoiface"
)

// UserCommand is a command paired with the return path to its sender
type UserCommand struct {
	User    User
	Command pbi.MessageV1
}

// TryReject sends a rejection event for the contained command unless it would block. Returns true when sent.
func (uc UserCommand) TryReject(err error) bool {
	return uc.User.TrySend(rejectedCmd(uc.Command, err))
}
