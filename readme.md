# learning-go

## Workspace

```bash
$GOPATH/go # Workspace
$GOPATH/go/src # Source code
$GOPATH/go/bin # Compiled binaries
```

## Commands

```bash
## Build and run
go run main.go
go build main.go #creates main.exe
go build -o binary main.go # creates binary.exe

## Format
## go prefers tabs for indentation and One True Brace Style over K&S

### 1. go fmt
go fmt main.go # provided by go

### 2. goimports
go install golang.org/x/tools/cmd/goimports@latest #install goimports for format import statements
goimports -l -w .

### 3.golint
go install golang.org/x/lint/golint@latest
golint ./...

### 4. go vet
go vet ./...

### 5. golangci-lint (golint + go vet)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run

go env # List all go env variables

## Install or update
go install github.com/rakyll/hey@latest

```
