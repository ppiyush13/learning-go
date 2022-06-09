# learning-go

## apt-get commands

```bash
# Add

# Remove
sudo apt-get remove --purge golang-go
sudo apt-get remove --purge go
sudo apt-get autoremove

# Search
sudo apt search golang-go

# Repository
sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
sudo add-apt-repository -r ppa:ubuntu-lxc/lxc-stable
```

## Workspace

```bash
$GOPATH/go # Workspace
$GOPATH/go/src # Source code
$GOPATH/go/bin # Compiled binaries
```

## Formatting

```bash
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
```

## Commands

```bash
## Build and run
go run main.go
go build main.go #creates main.exe
go build -o binary main.go # creates binary.exe

go env # List all go env variables

## Install or update
go install github.com/rakyll/hey@latest

```

## Secondary go

```bash
go get golang.org/dl/go.1.15.6
go1.15.6 download

go1.15.6 build # use this instead of go

go1.15.6 env GOROOT # get the root directory path
rm -rf $(go1.15.6 env GOROOT)
rm $(go env GOPATH)/bin/go1.15.6
```
