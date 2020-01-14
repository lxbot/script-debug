.PHONY: build

build:
	go build -buildmode=plugin -o script-debug.so script.go
