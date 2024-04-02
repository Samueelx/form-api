build:
	@go build -o ./bin/form ./cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/form
