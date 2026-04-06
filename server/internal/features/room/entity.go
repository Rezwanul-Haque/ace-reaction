package room

import "time"

type RoomStatus string

const (
	RoomStatusWaiting  RoomStatus = "waiting"
	RoomStatusPlaying  RoomStatus = "playing"
	RoomStatusFinished RoomStatus = "finished"
)

type Room struct {
	ID        string     `json:"id"`
	Players   []string   `json:"players"`
	Status    RoomStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

func NewRoom(id string) *Room {
	return &Room{
		ID:        id,
		Players:   make([]string, 0, 2),
		Status:    RoomStatusWaiting,
		CreatedAt: time.Now(),
	}
}

func (r *Room) AddPlayer(name string) bool {
	if len(r.Players) >= 2 {
		return false
	}
	for _, p := range r.Players {
		if p == name {
			return false
		}
	}
	r.Players = append(r.Players, name)
	if len(r.Players) == 2 {
		r.Status = RoomStatusPlaying
	}
	return true
}

func (r *Room) RemovePlayer(name string) {
	for i, p := range r.Players {
		if p == name {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			break
		}
	}
	if r.Status == RoomStatusPlaying {
		r.Status = RoomStatusFinished
	}
}

func (r *Room) IsFull() bool {
	return len(r.Players) >= 2
}

type RoomRepository interface {
	Create(room *Room) error
	FindByID(id string) (*Room, error)
	Delete(id string) error
}
