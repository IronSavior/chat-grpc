package chat

import (
	"fmt"

	"github.com/IronSavior/chat-grpc/chatpb/commands"
	"github.com/IronSavior/chat-grpc/chatpb/events"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/status"
	pbi "google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// mustAny attempts to encode the given message and panics on failure
func mustAny(m pbi.MessageV1) *anypb.Any {
	any, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("BUG: Cannot MarshalAny(%T): %s", m, err))
	}

	return any
}

func rejectedCmd(cmd pbi.MessageV1, err error) *events.Rejected {
	assertIsCommand(cmd)

	return &events.Rejected{
		Code:    uint32(status.Code(err)),
		Message: status.Convert(err).Message(),
		Command: mustAny(cmd),
	}
}

func unjoinCmd(err error) *commands.Unjoin {
	msg := "Leaving"
	if err != nil {
		msg = status.Convert(err).Message()
	}

	return &commands.Unjoin{
		Message: msg,
		Code:    uint32(status.Code(err)),
	}
}

func unjoinedEvt(u User, err error) *events.Unjoined {
	msg := "Leaving"
	if err != nil {
		msg = status.Convert(err).Message()
	}

	return &events.Unjoined{
		Time:     timestamppb.Now(),
		Code:     uint32(status.Code(err)),
		UserName: u.Name(),
		Message:  msg,
	}
}

func assertIsCommand(cmd pbi.MessageV1) {
	if !IsCommand(cmd) {
		panic(fmt.Sprintf("BUG: unknown event type: %T", cmd))
	}
}

func IsCommand(cmd pbi.MessageV1) bool {
	switch cmd.(type) {
	case *commands.Say:
	case *commands.Create:
	case *commands.Destroy:
	case *commands.Join:
	case *commands.Unjoin:
	default:
		return false
	}

	return true
}

// func assertIsEvent(event pbi.MessageV1) {
// 	if !IsEvent(event) {
// 		panic(fmt.Sprintf("BUG: unknown event type: %T", event))
// 	}
// }

func IsEvent(event pbi.MessageV1) bool {
	switch event.(type) {
	case *events.Said:
	case *events.Created:
	case *events.Destroyed:
	case *events.Joined:
	case *events.Unjoined:
	case *events.Rejected:
	default:
		return false
	}

	return true
}
