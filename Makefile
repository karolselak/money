all: build start
build:
	go build ./cmd/money

start:
	go run ./cmd/money
