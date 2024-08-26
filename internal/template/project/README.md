## For project
> contains the specific project files and folders.

1. generate the .gitignore file to ignore the specific files and folders.
```gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
tmp/

# IDE specific files
.vscode
.idea

# .env file
.env

# Project build
main
*templ.go
```

2. generate the Makefile file to automate the project tasks.
```makefile
.PHONY: build clean tool lint help

all: build

build:
    go build -v .

tool:
    go tool vet . |& grep -v vendor; true
    gofmt -w .

lint:
    golint ./...

clean:
    rm -rf go-gin-example
    go clean -i .

help:
    @echo "make: compile packages and dependencies"
    @echo "make tool: run specified go tool"
    @echo "make lint: golint ./..."
    @echo "make clean: remove object files and cached files"
```

3. generate the README.md file to describe the project.