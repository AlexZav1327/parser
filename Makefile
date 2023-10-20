build:
	go build -o ./bin/parser ./cmd/parser/main.go

fmt:
	gofumpt -w .

tidy:	
	go mod tidy

lint: build fmt tidy
	golangci-lint run ./...

run:
	go run ./cmd/parser/main.go