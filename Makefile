.PHONY: build
build:
	go build -o ./build/tasker ./cmd/tasker/main.go

run:
	go run ./cmd/tasker/main.go

start:
	./build/tasker
