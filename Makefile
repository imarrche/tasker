.PHONY: build
build:
	go build -o ./build/tasker ./cmd/tasker/main.go

run:
	go run ./cmd/tasker/main.go

start:
	./build/tasker

test:
	go test -cover -coverprofile=coverage.html -timeout 30s ./...

.PHONY: coverage
coverage:
	go tool cover -html=coverage.html
