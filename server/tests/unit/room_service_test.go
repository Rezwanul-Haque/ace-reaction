package unit_test

import (
	"testing"

	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/room"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	r, err := svc.CreateRoom()
	require.NoError(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, 8, len(r.ID))
}

func TestGetRoom(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	created, _ := svc.CreateRoom()
	found, err := svc.GetRoom(created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestGetRoomNotFound(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	_, err := svc.GetRoom("nonexistent")
	assert.Error(t, err)
}

func TestJoinRoom(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	created, _ := svc.CreateRoom()
	r, err := svc.JoinRoom(created.ID, "Alice")
	require.NoError(t, err)
	assert.Equal(t, 1, len(r.Players))
}

func TestJoinRoomFull(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	created, _ := svc.CreateRoom()
	svc.JoinRoom(created.ID, "Alice")
	svc.JoinRoom(created.ID, "Bob")

	_, err := svc.JoinRoom(created.ID, "Charlie")
	assert.Error(t, err)
}

func TestRemovePlayerDeletesEmptyRoom(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	svc := room.NewRoomService(repo)

	created, _ := svc.CreateRoom()
	svc.JoinRoom(created.ID, "Alice")

	svc.RemovePlayer(created.ID, "Alice")

	_, err := svc.GetRoom(created.ID)
	assert.Error(t, err)
}
