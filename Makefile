build:
	@go build -o bin/ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecommerce