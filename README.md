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

### Commands about module system 
- go list -m -versions [module] : go list shows used packages. -m flag make it show used module instead of pkgs. and -version flag module's available versions.
- go get [repository] : updates dependency of module. ( attach like @v1.2.0 for specific version) ( -u flag for update. -u=patch allows only patch updates)
- go mod tiny : removes non-using versions.
- go mod vendor : stores all dependencies in folder named vendor. If you wanna keep it up to date, you have to use this command everytime you change any dependency.

- Makefile ( make, make fmt, make build . . . )

  The code below is basic template of makefile for Go. the indent must be tab.

```make
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	shadow ./... # this tool detects shadowing variables
.PHONY:vet

build: vet
	go build
.PHONY:build
```

# Go tips
* pkg.go.dev
this service automatically indexes every open source Go projects.
