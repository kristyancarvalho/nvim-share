.PHONY: build run clean

build:
	go mod tidy
	go build -o bin/nvim-share cmd/main.go

run: build
	bin/./nvim-share

clean:
	rm -rf bin
