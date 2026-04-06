package game

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/rezwanul-haque/reflex-card-game/server/internal/infra"
)

const (
	RoundsToWin      = 3
	CardFlipMinDelay = 1500 * time.Millisecond
	CardFlipMaxDelay = 3000 * time.Millisecond
	AceClickTimeout  = 1000 * time.Millisecond
	RoundEndDelay    = 2000 * time.Millisecond
)

type GameRoom struct {
	Game        *Game
	Connections map[string]*infra.Connection
	StopChan    chan struct{}
}

type GameService struct {
	mu    sync.RWMutex
	rooms map[string]*GameRoom
}

func NewGameService() *GameService {
	return &GameService{
		rooms: make(map[string]*GameRoom),
	}
}

func (s *GameService) CreateGame(roomID, player1, player2 string, conns map[string]*infra.Connection) {
	g := NewGame(roomID, player1, player2, RoundsToWin)
	gameRoom := &GameRoom{
		Game:        g,
		Connections: conns,
		StopChan:    make(chan struct{}),
	}

	s.mu.Lock()
	s.rooms[roomID] = gameRoom
	s.mu.Unlock()

	// Notify players
	for _, playerName := range g.Players {
		conn := conns[playerName]
		opponent := g.GetOpponent(playerName)
		playerNum := 1
		if playerName == player2 {
			playerNum = 2
		}
		conn.SendJSON(infra.GameStartMsg{
			Type:         "game_start",
			Opponent:     opponent,
			PlayerNumber: playerNum,
		})
	}

	go s.runGameLoop(roomID)
}

func (s *GameService) GetGameRoom(roomID string) *GameRoom {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rooms[roomID]
}

func (s *GameService) HandleClick(roomID, player string) {
	s.mu.RLock()
	gameRoom := s.rooms[roomID]
	s.mu.RUnlock()

	if gameRoom == nil {
		return
	}

	g := gameRoom.Game
	g.Lock()
	defer g.Unlock()

	result := g.HandleClick(player)
	if result != nil {
		s.broadcastRoundResult(gameRoom, result)
	}
}

func (s *GameService) HandleDisconnect(roomID, player string) {
	s.mu.Lock()
	gameRoom := s.rooms[roomID]
	s.mu.Unlock()

	if gameRoom == nil {
		return
	}

	// Stop game loop
	select {
	case <-gameRoom.StopChan:
	default:
		close(gameRoom.StopChan)
	}

	// Notify opponent
	g := gameRoom.Game
	opponent := g.GetOpponent(player)
	if conn, ok := gameRoom.Connections[opponent]; ok {
		conn.SendJSON(infra.PlayerLeftMsg{
			Type:   "player_left",
			Player: player,
		})
		// Award win to remaining player
		conn.SendJSON(infra.GameOverMsg{
			Type:   "game_over",
			Winner: opponent,
			Scores: g.Scores,
		})
	}

	s.mu.Lock()
	delete(s.rooms, roomID)
	s.mu.Unlock()
}

func (s *GameService) runGameLoop(roomID string) {
	s.mu.RLock()
	gameRoom := s.rooms[roomID]
	s.mu.RUnlock()

	if gameRoom == nil {
		return
	}

	// Initial delay before first card
	select {
	case <-time.After(2 * time.Second):
	case <-gameRoom.StopChan:
		return
	}

	for {
		select {
		case <-gameRoom.StopChan:
			return
		default:
		}

		g := gameRoom.Game
		g.Lock()

		if g.State == GameStateFinished {
			g.Unlock()
			return
		}

		card, cardNum, ok := g.FlipCard()
		g.Unlock()

		if !ok {
			return
		}

		// Broadcast card flip
		s.broadcast(gameRoom, infra.CardFlipMsg{
			Type:       "card_flip",
			Card:       card,
			CardNumber: cardNum,
		})

		if card.IsAce() {
			// Wait for clicks with timeout
			select {
			case <-time.After(AceClickTimeout):
			case <-gameRoom.StopChan:
				return
			}

			// Resolve if only one player clicked
			g.Lock()
			result := g.ResolveAceTimeout()
			if result != nil {
				g.Unlock()
				s.broadcastRoundResult(gameRoom, result)
			} else if !anyClicked(g) {
				// Nobody clicked — just continue
				g.Unlock()
			} else {
				g.Unlock()
			}
		} else {
			// Non-ace card: wait for potential early click or timeout
			delay := CardFlipMinDelay + time.Duration(rand.Int63n(int64(CardFlipMaxDelay-CardFlipMinDelay)))
			select {
			case <-time.After(delay):
			case <-gameRoom.StopChan:
				return
			}
		}

		// Check if round ended (from a click during the wait)
		g.Lock()
		if g.State == GameStateRoundEnd {
			// Check game over
			over, winner := g.IsGameOver()
			if over {
				g.Unlock()
				s.broadcast(gameRoom, infra.GameOverMsg{
					Type:   "game_over",
					Winner: winner,
					Scores: g.Scores,
				})
				s.cleanup(roomID)
				return
			}
			g.PrepareNextRound()
			g.Unlock()

			// Pause between rounds
			select {
			case <-time.After(RoundEndDelay):
			case <-gameRoom.StopChan:
				return
			}
		} else {
			g.Unlock()
		}
	}
}

func anyClicked(g *Game) bool {
	for _, p := range g.Players {
		if g.HasClicked[p] {
			return true
		}
	}
	return false
}

func (s *GameService) broadcastRoundResult(gameRoom *GameRoom, result *RoundResult) {
	s.broadcast(gameRoom, infra.RoundResultMsg{
		Type:   "round_result",
		Winner: result.Winner,
		Loser:  result.Loser,
		Reason: result.Reason,
		Scores: gameRoom.Game.Scores,
	})
}

func (s *GameService) broadcast(gameRoom *GameRoom, msg any) {
	for _, conn := range gameRoom.Connections {
		if err := conn.SendJSON(msg); err != nil {
			log.Printf("error broadcasting to %s: %v", conn.PlayerName, err)
		}
	}
}

func (s *GameService) cleanup(roomID string) {
	s.mu.Lock()
	gameRoom := s.rooms[roomID]
	if gameRoom != nil {
		select {
		case <-gameRoom.StopChan:
		default:
			close(gameRoom.StopChan)
		}
		delete(s.rooms, roomID)
	}
	s.mu.Unlock()
}
