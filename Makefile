## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build:
	@echo "Building binary ..."
	env GOOS=linux CGO_ENABLED=0 go build -o main ./main.go
	@echo "Done!"

## build_up: stops docker-compose (if running), builds all projects and starts docker compose
build_up: down build
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## run: build & runs the binary (windows), this prevents windows firewall from showing a popup
run:
	go build -o main.exe ./main.go
	main.exe

## coverage: runs the tests and generates a coverage report
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out