build:
	@go build -o bin/tiago-ecom-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/tiago-ecom-api
