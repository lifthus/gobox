# Installing Go
```
// Mac
brew install go
```
```
// in $HOME/.profile or .zshrc
export GOPATH=$HOME
export PATH=$PATH:$GOPATH/bin
```

# Basic Go commands
- go fmt : formatting codes in standard format
- goimports -l -w . : sort import lines
- golint ./... : according to style guide, suggests proper var name, error message formatting, comment placement etc.
- go vet ./... : capturing unintentional but syntactically correct code
- golanci-lint run : integrating code quality tools ( with .golangci.yml file )
- go build -o [output] [target]
- go get [repository]

- Makefile ( make, make fmt, make build . . . )

  The code below is basic template of makefile for Go. the indent must be tab.

```
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build
.PHONY:build
```

# Go tips

- To save all dependencies at vendor folder, use the command below

  ```
  go mod vendor
  ```

  if you wanna keep it up to date, you have to use this command everytime you change the dependency.
