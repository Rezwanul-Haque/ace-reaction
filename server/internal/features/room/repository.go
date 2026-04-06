package room

import (
	"fmt"
	"sync"
)

type MemoryRoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]*Room
}

func NewMemoryRoomRepository() *MemoryRoomRepository {
	return &MemoryRoomRepository{
		rooms: make(map[string]*Room),
	}
}

func (r *MemoryRoomRepository) Create(room *Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.rooms[room.ID]; exists {
		return fmt.Errorf("room %s already exists", room.ID)
	}
	r.rooms[room.ID] = room
	return nil
}

func (r *MemoryRoomRepository) FindByID(id string) (*Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	room, exists := r.rooms[id]
	if !exists {
		return nil, fmt.Errorf("room %s not found", id)
	}
	return room, nil
}

func (r *MemoryRoomRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.rooms, id)
	return nil
}
