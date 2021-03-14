package chat

import (
	pbi "google.golang.org/protobuf/runtime/protoiface"
)

// Stream represents a connected chat client and is a combined command input path and event return path.
// NB: Implementations MUST be comparable--such as a pointer to a struct
type Stream interface {
	CommandStream
	EventStream
}

// CommandStream is a source of chat commands
type CommandStream interface {
	// Channel from which the service may receive commands. Must be closed by impl to signal end.
	Commands() <-chan pbi.MessageV1

	// Returns the error state (only valid after the Commands channel has closed)
	Err() error
}

// EventStream is an event sink
type EventStream interface {
	// Channel for the service to send chat events. Will not be directly closed by
	// service. CloseEvents() may close it.
	Events() chan<- pbi.MessageV1

	// Signals the termination of the client session by the service with optional error. No further events will be sent
	// and any commands will be rejected.
	CloseEvents(error)
}

// Events composes additional behavior on an EventStream
type Events struct {
	EventStream
}

// TrySend sends the event on the stream unless it would block. Returns true when event is sent.
func (e Events) TrySend(event pbi.MessageV1) bool {
	select {
	case e.Events() <- event:
		return true
	default:
		return false
	}
}
