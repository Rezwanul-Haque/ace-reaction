.PHONY: dev dev-down up down build test test-server test-client clean

# Development (hot reload)
dev:
	docker compose up --build

dev-down:
	docker compose down

# Production
up:
	docker compose -f docker-compose.prod.yml up --build -d

down:
	docker compose -f docker-compose.prod.yml down

# Build
build:
	docker compose -f docker-compose.prod.yml build

# Tests
test: test-server

test-server:
	cd server && go test ./... -v

test-client:
	cd client && npm test

# Cleanup
clean:
	docker compose down -v --rmi local
	docker compose -f docker-compose.prod.yml down -v --rmi local
