.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:golint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build .
.PHONY: build
