package chat

import (
	"sync"

	pbi "google.golang.org/protobuf/runtime/protoiface"
)

type TStream struct {
	events   chan pbi.MessageV1
	commands chan pbi.MessageV1

	err    error
	close1 sync.Once
}

func NewTStream(qSize int) *TStream {
	return &TStream{
		events:   make(chan pbi.MessageV1, qSize),
		commands: make(chan pbi.MessageV1, qSize),
	}
}

func (ts *TStream) Err() error {
	return ts.err
}

func (ts *TStream) Commands() <-chan pbi.MessageV1 {
	return ts.commands
}

func (ts *TStream) TryCommand(cmd pbi.MessageV1) bool {
	if !IsCommand(cmd) {
		return false
	}

	select {
	case ts.commands <- cmd:
		return true
	default:
		return false
	}
}

func (ts *TStream) Events() chan<- pbi.MessageV1 {
	return ts.events
}

func (ts *TStream) CloseEvents(err error) {
	ts.close1.Do(func() {
		ts.err = err
		close(ts.events)
	})
}
