run: build
	@./bin/main

build: $(wildcard *.go)
	@go build -o bin/main main.go
