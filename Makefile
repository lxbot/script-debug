.PHONY: build debug

build:
	go build -buildmode=plugin -o script-debug.so script.go

debug:
	go build -gcflags="all=-N -l" -buildmode=plugin -o script-debug.so script.go
