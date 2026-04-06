package room

import (
	"fmt"

	"github.com/google/uuid"
)

type RoomService struct {
	repo RoomRepository
}

func NewRoomService(repo RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom() (*Room, error) {
	id := uuid.New().String()[:8]
	room := NewRoom(id)
	if err := s.repo.Create(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomService) GetRoom(id string) (*Room, error) {
	return s.repo.FindByID(id)
}

func (s *RoomService) JoinRoom(id, playerName string) (*Room, error) {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if !room.AddPlayer(playerName) {
		return nil, fmt.Errorf("cannot join room: room is full or player name taken")
	}
	return room, nil
}

func (s *RoomService) RemovePlayer(id, playerName string) {
	room, err := s.repo.FindByID(id)
	if err != nil {
		return
	}
	room.RemovePlayer(playerName)
	if len(room.Players) == 0 {
		s.repo.Delete(id)
	}
}
