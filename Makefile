.DEFAULT_GOAL := help
.PHONY: help files build

help:
	@echo "Usage: make [TARGET]"
	@echo "Targets:"
	@echo "files           Show files"
	@echo "build           Build cli"

files:
	@find . -path './.git' -prune -o -ls > FILES

build:
	@go mod tidy && CGO_ENABLED="0" go build -o lddc app/*.go
