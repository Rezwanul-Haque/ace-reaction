package unit_test

import (
	"testing"

	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/room"
	"github.com/stretchr/testify/assert"
)

func TestNewRoom(t *testing.T) {
	r := room.NewRoom("test-room")
	assert.Equal(t, "test-room", r.ID)
	assert.Equal(t, room.RoomStatusWaiting, r.Status)
	assert.Empty(t, r.Players)
	assert.False(t, r.IsFull())
}

func TestAddFirstPlayer(t *testing.T) {
	r := room.NewRoom("test-room")
	ok := r.AddPlayer("Alice")
	assert.True(t, ok)
	assert.Equal(t, 1, len(r.Players))
	assert.Equal(t, room.RoomStatusWaiting, r.Status)
	assert.False(t, r.IsFull())
}

func TestAddSecondPlayer(t *testing.T) {
	r := room.NewRoom("test-room")
	r.AddPlayer("Alice")
	ok := r.AddPlayer("Bob")
	assert.True(t, ok)
	assert.Equal(t, 2, len(r.Players))
	assert.Equal(t, room.RoomStatusPlaying, r.Status)
	assert.True(t, r.IsFull())
}

func TestAddThirdPlayerFails(t *testing.T) {
	r := room.NewRoom("test-room")
	r.AddPlayer("Alice")
	r.AddPlayer("Bob")
	ok := r.AddPlayer("Charlie")
	assert.False(t, ok)
	assert.Equal(t, 2, len(r.Players))
}

func TestAddDuplicatePlayerFails(t *testing.T) {
	r := room.NewRoom("test-room")
	r.AddPlayer("Alice")
	ok := r.AddPlayer("Alice")
	assert.False(t, ok)
	assert.Equal(t, 1, len(r.Players))
}

func TestRemovePlayer(t *testing.T) {
	r := room.NewRoom("test-room")
	r.AddPlayer("Alice")
	r.AddPlayer("Bob")

	r.RemovePlayer("Alice")
	assert.Equal(t, 1, len(r.Players))
	assert.Equal(t, "Bob", r.Players[0])
	assert.Equal(t, room.RoomStatusFinished, r.Status)
}

func TestRemoveNonexistentPlayer(t *testing.T) {
	r := room.NewRoom("test-room")
	r.AddPlayer("Alice")
	r.RemovePlayer("Charlie") // should not panic
	assert.Equal(t, 1, len(r.Players))
}
