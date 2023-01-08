# Go tips

- To save all dependencies at vendor folder, use the command below

  ```
  go mod vendor
  ```

  if you wanna keep it up to date, you have to use this command everytime you change the dependency.

- Makefile
  The code below is basic template of makefile

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
