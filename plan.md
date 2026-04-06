# Reflex Card Game — Implementation Plan

## Problem Summary

A real-time two-player reflex card game — cards flip one at a time, players race to click when an Ace appears. Click too early (non-Ace) and you lose. First to click on an Ace wins. Must work over the internet with a live deployment.

---

## Tech Stack

### Backend: Go + Echo Framework

| Reason | Detail |
|--------|--------|
| Echo framework | Lightweight, fast, built-in WebSocket support, middleware ecosystem |
| Vertical slice + clean arch | Feature-organized code with clear domain/service/handler layers per slice |
| Concurrency model | Goroutines + channels for game rooms, player connections, card timers |
| Low latency | Compiled binary, minimal GC pauses — critical for a reflex game |
| Simple deployment | Single static binary, easy to containerize |

### Frontend: React + TypeScript + Vite + TanStack Query

- React + TypeScript for component-based UI with type safety
- Vite for fast dev/build
- **TanStack Query** for all REST API calls (room creation, game state)
- Native `WebSocket` API for real-time game events
- Tailwind CSS for styling
- Framer Motion for card flip animations

### Infrastructure

- Docker + docker-compose for local development
- Makefile: `make dev` to start, `make dev-down` to stop
- Multi-stage Dockerfile for production deployment

---

## Architecture

### High-Level

```
Browser A ──WebSocket──┐
                       ├── Go Echo Server (vertical slices)
Browser B ──WebSocket──┘
                       └── serves React static files
```

All game logic lives on the server. Clients are thin — they render what the server tells them and send click events.

### Backend — Vertical Slice + Clean Architecture

Each feature is a self-contained vertical slice with clean architecture layers:

```
server/
├── internal/
│   ├── room/                    # Room slice
│   │   ├── domain/
│   │   │   └── room.go          # Room entity, interfaces
│   │   ├── service/
│   │   │   └── room_service.go  # Business logic
│   │   ├── handler/
│   │   │   └── room_handler.go  # Echo HTTP handlers
│   │   └── repository/
│   │       └── memory_repo.go   # In-memory store
│   │
│   ├── game/                    # Game slice
│   │   ├── domain/
│   │   │   ├── card.go          # Card, Deck entities
│   │   │   ├── game.go          # Game entity, state machine
│   │   │   └── interfaces.go    # Repository/service interfaces
│   │   ├── service/
│   │   │   └── game_service.go  # Game engine logic
│   │   ├── handler/
│   │   │   └── ws_handler.go    # WebSocket handler
│   │   └── repository/
│   │       └── memory_repo.go   # In-memory game store
│   │
│   └── shared/                  # Cross-cutting concerns
│       ├── ws/
│       │   └── connection.go    # WebSocket connection wrapper
│       └── config/
│           └── config.go        # App configuration
│
├── main.go                      # Entry point, wire up slices
├── go.mod
└── go.sum
```

**Why vertical slices:** Each feature (room, game) is independently understandable. You can read one folder and know everything about that feature — no jumping across layers scattered in different directories.

**Why clean architecture inside each slice:** Domain logic stays pure (no framework imports), services orchestrate business rules, handlers deal with HTTP/WS concerns, repositories abstract storage. Easy to test each layer independently.

### Frontend Structure

```
client/
├── src/
│   ├── api/
│   │   ├── client.ts            # Axios/fetch base config
│   │   └── rooms.ts             # TanStack Query hooks for room API
│   ├── components/
│   │   ├── Card.tsx
│   │   ├── Scoreboard.tsx
│   │   └── PlayerBadge.tsx
│   ├── pages/
│   │   ├── Home.tsx
│   │   ├── Lobby.tsx
│   │   ├── Game.tsx
│   │   └── GameOver.tsx
│   ├── hooks/
│   │   ├── useGameWebSocket.ts  # WebSocket connection + message handling
│   │   └── useGameState.ts      # Game state derived from WS messages
│   ├── types/
│   │   └── game.ts              # Shared type definitions
│   ├── App.tsx
│   └── main.tsx
├── package.json
├── tsconfig.json
├── tailwind.config.js
└── vite.config.ts
```

---

## WebSocket Message Protocol

```json
// Client → Server
{ "type": "click" }

// Server → Client
{ "type": "waiting",       "room_id": "abc123" }
{ "type": "game_start",    "opponent": "Bob", "player_number": 1 }
{ "type": "card_flip",     "card": { "suit": "hearts", "rank": "7" }, "card_number": 3 }
{ "type": "round_result",  "winner": "Alice", "reason": "ace_click" }
{ "type": "round_result",  "winner": "Bob",   "reason": "early_click", "loser": "Alice" }
{ "type": "game_over",     "winner": "Alice",  "score": { "Alice": 3, "Bob": 1 } }
{ "type": "player_left",   "player": "Bob" }
{ "type": "error",         "message": "room is full" }
```

## REST API

```
POST   /api/rooms              → { "room_id": "abc123" }
GET    /api/rooms/:id          → { "room_id": "abc123", "players": 1, "status": "waiting" }
GET    /ws?room=abc123&name=Alice  → WebSocket upgrade
```

TanStack Query handles the REST calls; WebSocket handles all real-time game events.

---

## Task Breakdown

### Task 1 — Project scaffolding

- Init Go module with Echo, set up vertical slice folder structure
- Scaffold React app with Vite + TypeScript
- Add Tailwind CSS, TanStack Query
- Create Dockerfile, docker-compose.yml, Makefile
- Set up `.gitignore`, CLAUDE.md

**Why:** Clean structure from the start. `make dev` works from day one.

---

### Task 2 — Room slice (Backend)

- `Room` domain entity: ID, players, status (waiting/playing/finished)
- `RoomRepository` interface + in-memory implementation
- `RoomService`: create room, join room, get room status
- `RoomHandler`: `POST /api/rooms`, `GET /api/rooms/:id`
- Unit tests for service and repository

**Why:** Rooms are the foundation. REST endpoints let the frontend create/query rooms via TanStack Query before upgrading to WebSocket.

---

### Task 3 — Game slice — domain & engine (Backend)

- `Card` and `Deck` domain entities (52 cards, shuffle)
- `Game` domain entity: state machine (waiting → playing → round_end → finished)
- `GameService`: start game, flip card, handle click, determine winner
- Click validation: Ace + click = win, non-Ace + click = lose, tie-breaking by server timestamp
- Round tracking (best of 5)
- Unit tests for game engine logic

**Why:** Server-authoritative game logic is the core of the project. Must be thoroughly tested.

---

### Task 4 — WebSocket handler (Backend)

- Echo WebSocket upgrade endpoint
- Connection wrapper with read/write pumps
- Wire WebSocket events to GameService
- Per-room message broadcasting
- Player disconnect handling
- Ping/pong keepalive

**Why:** Bridges the real-time communication layer to the game engine.

---

### Task 5 — Frontend: home & lobby (TanStack Query + WebSocket)

- Home page: "Create Game" / "Join Game"
- `useCreateRoom` mutation (TanStack Query) → `POST /api/rooms`
- `useRoomStatus` query (TanStack Query) → `GET /api/rooms/:id`
- Lobby/waiting screen with room code to share
- WebSocket connection established on join
- Transition to game when both players connect

**Why:** TanStack Query handles REST API state cleanly. WebSocket takes over for real-time game events.

---

### Task 6 — Frontend: game screen

- Card display with flip animation (Framer Motion)
- "SLAP!" button — sends `click` via WebSocket
- Scoreboard component
- Round result overlay
- Game over screen with final scores
- "Play Again" option

**Why:** The interactive game UI — this is what makes it feel like a product.

---

### Task 7 — Edge cases & polish

- Player disconnect → opponent wins by forfeit
- Prevent clicking when no card is shown
- Disable button briefly after round ends
- Mobile-friendly responsive layout
- Loading/error states
- Visual card suits (red/black colors)

**Why:** Production-quality polish as the task requires.

---

### Task 8 — Unit & E2E tests

- **Unit tests (Go):**
  - Game engine: deck shuffle, card flip, click validation, scoring
  - Room service: create, join, status transitions
  - Message serialization
- **E2E tests:**
  - Two players join a room
  - Full game flow: cards flip, player clicks on Ace, wins round
  - Early click penalty
  - Player disconnect handling
- Test tooling: Go `testing` package, `testify` for assertions

**Why:** Tests prove the game logic is correct and the real-time flow works end-to-end.

---

### Task 9 — Docker, Compose & Makefile

- `Dockerfile`: multi-stage (build React → build Go → minimal runtime image)
- `docker-compose.yml`: backend + frontend services for development with hot reload
- `Makefile`:
  - `make dev` — `docker-compose up --build`
  - `make dev-down` — `docker-compose down`
  - `make test` — run all tests
  - `make build` — production build

**Why:** Single command to run the entire project. No manual setup.

---

### Task 10 — Deployment & README

- Deploy to Fly.io or Railway
- Write README: architecture, tech decisions, tradeoffs, run instructions
- Include architecture diagram

**Why:** Explicit deliverable. README explains the "why" behind decisions.

---

## Key Design Decisions

1. **Vertical slice architecture** — Features are self-contained. Each slice (room, game) has its own domain/service/handler/repository. No cross-slice coupling.
2. **Clean architecture within slices** — Domain layer has zero external dependencies. Business logic is testable without HTTP or WebSocket concerns.
3. **Server-authoritative** — All game state and timing controlled by the server. Clients are dumb renderers.
4. **TanStack Query for REST, raw WebSocket for real-time** — TanStack Query gives caching, loading states, and error handling for free on REST calls. WebSocket handles the latency-critical game events.
5. **Echo framework** — Lightweight, fast, built-in middleware and WebSocket support. No unnecessary abstractions.
6. **No database** — In-memory repositories. Games are ephemeral. Repository interfaces make it trivial to swap in Redis/DB later.
7. **Docker Compose for dev** — `make dev` and you're running. No "install Go, install Node, run two terminals" friction.

---

## Tradeoffs

- **Network latency** — A player with lower latency has a slight advantage. Would need latency compensation for true fairness (out of scope).
- **No persistence** — Server restart loses all games. Acceptable for demo; repository interfaces allow easy Redis/DB addition.
- **No auth** — Players identified by name only. Would add JWT/session tokens for production.
- **Single server** — No horizontal scaling. Would need shared state store for multi-instance.
- **In-memory repository** — Simple but not production-grade. Clean architecture makes swapping trivial.
