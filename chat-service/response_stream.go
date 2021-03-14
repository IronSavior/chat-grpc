package chat

import (
	"sync"

	pbi "google.golang.org/protobuf/runtime/protoiface"
)

// ResponseStream is a simple EventStream for answering syncronous one-shot calls
type ResponseStream struct {
	events chan pbi.MessageV1

	err    error
	close1 sync.Once
}

// Channel for the service to send chat events. Will not be directly closed by
// service. CloseEvents() may close it.
func (rs *ResponseStream) Events() chan<- pbi.MessageV1 {
	return rs.events
}

// CloseEvents records the given error and ends the event channel
func (rs *ResponseStream) CloseEvents(err error) {
	rs.close1.Do(func() {
		rs.err = err
		close(rs.events)
	})
}

func (rs *ResponseStream) Err() error {
	return rs.err
}
