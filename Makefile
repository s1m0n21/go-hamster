SHELL=/usr/bin/env bash

GO_VERSION:=$(shell go version | cut -d' ' -f 3 | cut -d. -f 2)
ifeq ($(shell expr $(GO_VERSION) \< 13), 1)
$(warning Your Golang version is go 1.$(GO_VERSION))
$(error Update Golang to version $(shell grep '^go' go.mod))
endif

install:
	go build -o go-hamster ./cmd/go-hamster
	mv ./go-hamster /usr/local/bin

clean:
	rm -f /usr/local/bin/go-hamster