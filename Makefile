.PHONY: migrate createdb migrate_create force version migrate_up migrate_down

# ==============================================================================
# Go migrate postgresql

migrate_create:
	migrate create -seq -ext=.sql -dir=./migrations $(name)

force:
	migrate -database postgres://postgres:postgres@localhost:5432/go_store?sslmode=disable -path migrations force 1

version:
	migrate -database postgres://postgres:postgres@localhost:5432/go_store?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/go_store?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/go_store?sslmode=disable -path migrations down 1

# ==============================================================================
# Docker compose command

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go
	
build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ================================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)