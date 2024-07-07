build:
	@go build -o ./bin/rollin-my-own ./cmd/main/main.go

run:
	@go run ./cmd/main/main.go

test:
	@go test -v

lint: 
	@golangci-lint run ./...