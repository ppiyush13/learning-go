.DEFAULT_GOAL := run

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: run
run: fmt
	go run main.go
