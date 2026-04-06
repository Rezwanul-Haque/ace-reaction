package unit_test

import (
	"testing"

	"github.com/rezwanul-haque/reflex-card-game/server/internal/features/room"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAndFind(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	r := room.NewRoom("test-room")

	err := repo.Create(r)
	require.NoError(t, err)

	found, err := repo.FindByID("test-room")
	require.NoError(t, err)
	assert.Equal(t, r.ID, found.ID)
}

func TestCreateDuplicate(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	r := room.NewRoom("test-room")

	err := repo.Create(r)
	require.NoError(t, err)

	err = repo.Create(r)
	assert.Error(t, err)
}

func TestFindNonexistent(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	_, err := repo.FindByID("nonexistent")
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	repo := room.NewMemoryRoomRepository()
	r := room.NewRoom("test-room")
	repo.Create(r)

	err := repo.Delete("test-room")
	require.NoError(t, err)

	_, err = repo.FindByID("test-room")
	assert.Error(t, err)
}
