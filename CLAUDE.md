# Project: Reflex Card Game

## Rules
- **NEVER commit without explicit user permission**
- **For any new API, write both unit tests and e2e tests**
- Use **vertical slice architecture** + **clean architecture** on the backend
- Use **Golang Echo framework** for the backend
- Use **TanStack Query** (React Query) for API calling on the frontend
- Use **React + TypeScript + Vite** for the frontend
- Use **Tailwind CSS** for styling
- Run the whole project with `make dev` and tear down with `make dev-down`
- Docker + docker-compose for local development

## Architecture
- Backend: Vertical slices organized by feature (game, room, player), each slice contains handler, service, domain, repository layers following clean architecture
- Frontend: React + TypeScript + Vite + TanStack Query + Tailwind CSS
- Communication: WebSockets for real-time game state, REST for room creation/joining
- Server-authoritative game logic

## Tech Stack
- Go + Echo (backend)
- React + TypeScript + Vite (frontend)
- TanStack Query (API calls)
- Tailwind CSS (styling)
- Docker + docker-compose (containerization)
- Makefile (dev commands)
