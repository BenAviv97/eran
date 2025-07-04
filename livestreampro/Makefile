# /home/${USER}/livestreampro/Makefile
# /home/${USER}/livestreampro/Makefile

.PHONY: dev proto lint test-backend test-frontend pre-commit chaos-smoke budget-check lighthouse

dev:
	docker compose -f infra/docker/compose.local.yaml up --build

proto:
	buf generate ./backend/proto

lint:
	golangci-lint run ./...
	npm run --prefix frontend lint

chaos-smoke:
	@echo "Running chaos experiments"

budget-check:
	@echo "Checking Kubecost budgets"

lighthouse:
	npm run --prefix frontend lighthouse

test-backend:
	cd backend && go test ./...

test-frontend:
	npm run --prefix frontend test

pre-commit: lint test-backend

