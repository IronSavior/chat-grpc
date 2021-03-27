package chat

import (
	"testing"
	"time"

	"github.com/IronSavior/chat-grpc/chatpb/events"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	pbi "google.golang.org/protobuf/runtime/protoiface"
)

func AssertJoined(t *testing.T, event pbi.MessageV1, uName string) {
	t.Helper()

	require.IsType(t, (*events.Joined)(nil), event)
	joined := event.(*events.Joined)
	require.Equal(t, uName, joined.UserName)
	require.WithinDuration(t, time.Now(), joined.Time.AsTime(), 1*time.Second)
}

func AssertSaid(t *testing.T, event pbi.MessageV1, uName, msg string) {
	t.Helper()

	require.IsType(t, (*events.Said)(nil), event)
	said := event.(*events.Said)
	require.Equal(t, uName, said.UserName)
	require.Equal(t, msg, said.Message)
	require.WithinDuration(t, time.Now(), said.Time.AsTime(), 1*time.Second)
}

func AssertUnjoined(t *testing.T, event pbi.MessageV1, userName, msg string, code codes.Code) {
	t.Helper()

	require.IsType(t, (*events.Unjoined)(nil), event)
	unjoined := event.(*events.Unjoined)
	require.Equal(t, uint32(code), unjoined.Code)
	require.Equal(t, userName, unjoined.UserName)
	require.Equal(t, msg, unjoined.Message)
	require.WithinDuration(t, time.Now(), unjoined.Time.AsTime(), 1*time.Second)
}
