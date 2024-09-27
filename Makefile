.DEFAULT_GOAL := help
.PHONY: help build

help:
	@echo "Usage: make [TARGET]"
	@echo "Targets:"
	@echo "build           Build"

build:
	@go mod tidy && CGO_ENABLED="0" go build -o lddc app/*.go
