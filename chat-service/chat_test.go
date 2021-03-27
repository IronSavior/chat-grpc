package chat

import (
	"context"
	"testing"
	"time"

	"github.com/IronSavior/chat-grpc/chatpb/commands"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestChat(t *testing.T) {
	t.Parallel()

	t.Run("join, say, unjoin", func(t *testing.T) {
		t.Parallel()

		require := require.New(t)
		chat := NewChat("TestChat", 5)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		stream := NewTStream(5)
		go chat.Run(ctx)

		go chat.Stream(ctx, "TestUser", stream, &commands.Join{ChatName: "TestChat"})
		AssertJoined(t, <-stream.events, "TestUser")

		require.True(stream.TryCommand(&commands.Say{Message: "test message"}))
		AssertSaid(t, <-stream.events, "TestUser", "test message")

		require.True(stream.TryCommand(&commands.Unjoin{Message: "Leaving"}))
		AssertUnjoined(t, <-stream.events, "TestUser", "Leaving", codes.OK)
	})

	t.Run("join, say, unjoin (2 users)", func(t *testing.T) {
		t.Parallel()

		require := require.New(t)
		chat := NewChat("TestChat", 5)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		s1 := NewTStream(1)
		s2 := NewTStream(1)
		go chat.Run(ctx)

		go chat.Stream(ctx, "TestUser1", s1, &commands.Join{ChatName: "TestChat"})
		AssertJoined(t, <-s1.events, "TestUser1")

		go chat.Stream(ctx, "TestUser2", s2, &commands.Join{ChatName: "TestChat"})
		AssertJoined(t, <-s1.events, "TestUser2")
		AssertJoined(t, <-s2.events, "TestUser2")

		require.True(s1.TryCommand(&commands.Say{Message: "message 1"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 1")
		AssertSaid(t, <-s2.events, "TestUser1", "message 1")

		require.True(s2.TryCommand(&commands.Say{Message: "message 2"}))
		AssertSaid(t, <-s1.events, "TestUser2", "message 2")
		AssertSaid(t, <-s2.events, "TestUser2", "message 2")

		require.True(s1.TryCommand(&commands.Unjoin{Message: "Leaving"}))
		AssertUnjoined(t, <-s1.events, "TestUser1", "Leaving", codes.OK)
		AssertUnjoined(t, <-s2.events, "TestUser1", "Leaving", codes.OK)

		require.True(s2.TryCommand(&commands.Unjoin{Message: "Leaving"}))
		AssertUnjoined(t, <-s2.events, "TestUser2", "Leaving", codes.OK)
	})

	t.Run("stalled user is removed", func(t *testing.T) {
		t.Parallel()

		require := require.New(t)
		chat := NewChat("TestChat", 5)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		s1 := NewTStream(5)
		s2 := NewTStream(5) // stalled because we won't service the event queue after join
		go chat.Run(ctx)

		go chat.Stream(ctx, "TestUser1", s1, &commands.Join{ChatName: "TestChat"})
		AssertJoined(t, <-s1.events, "TestUser1")

		go chat.Stream(ctx, "TestUser2", s2, &commands.Join{ChatName: "TestChat"})
		AssertJoined(t, <-s1.events, "TestUser2")
		AssertJoined(t, <-s2.events, "TestUser2")

		require.True(s1.TryCommand(&commands.Say{Message: "message 1"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 1")

		require.True(s1.TryCommand(&commands.Say{Message: "message 2"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 2")

		require.True(s1.TryCommand(&commands.Say{Message: "message 3"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 3")

		require.True(s1.TryCommand(&commands.Say{Message: "message 4"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 4")

		require.True(s1.TryCommand(&commands.Say{Message: "message 5"}))
		AssertSaid(t, <-s1.events, "TestUser1", "message 5")

		require.True(s1.TryCommand(&commands.Say{Message: "message 6"})) // User 2 will stall here
		AssertSaid(t, <-s1.events, "TestUser1", "message 6")

		AssertUnjoined(t, <-s1.events, "TestUser2", "user unresponsive", codes.DataLoss)

		require.True(s1.TryCommand(&commands.Unjoin{Message: "Leaving"}))
		AssertUnjoined(t, <-s1.events, "TestUser1", "Leaving", codes.OK)
	})
}
