BINARY_NAME=gofun

all: build

build:
	go build -o $(BINARY_NAME) src/main.go

run: build
	./$(BINARY_NAME)

.PHONY: all build run

