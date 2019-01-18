#@Author: devuser@gmail.com
#@date: 2019-01-14
#@comment:
# Go parameters
BINARY_NAME=case_makefile
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_386=$(BINARY_NAME)_386_win32.exe
BINARY_AMD64=$(BINARY_NAME)_amd64_win64.exe

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOCMD=go
GOFMT ?= gofmt "-s"
GOGET=$(GVT) fetch
GOTEST=$(GOCMD) test
GVT=gvt
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
VETPACKAGES ?= $(shell go list ./... | grep -v /vendor/ | grep -v /examples/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
# GOPATH=/data/boyosoft/goBillProcess
all: test build
build:
    $(GOBUILD) -o $(BINARY_NAME) -v
test:
    $(GOTEST) -v ./...
clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)
    rm -f $(BINARY_linux)
    rm -f "case_makefile.log"
    rm -rf "test.db"
run: clean
    # $(GOBUILD) -o $(BINARY_NAME) -v ./...
    # ./$(BINARY_NAME)
    $(GOCMD) run main.go
deps:
    cd $(GOPATH)/src
    # $(GOGET) github.com/markbates/goth
    # $(GOGET) github.com/markbates/pop
    @hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
        gvt fetch github.com/golang/lint/golint; \
    fi
    @hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
        gvt fetch github.com/client9/misspell/cmd/misspell; \
    fi
    @hash govendor > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
        gvt fetch github.com/kardianos/govendor; \
    fi
    @hash embedmd > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
        gvt fetch github.com/campoy/embedmd; \
    fi
    cd $(GOPATH)/src/github.com/boyosoft/case_makefile

.PHONY: fmt
fmt:
    $(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
    # get all go files and run go fmt on them
    @diff=$$($(GOFMT) -d $(GOFILES)); \
    if [ -n "$$diff" ]; then \
        echo "Please run 'make fmt' and commit the result:"; \
        echo "$${diff}"; \
        exit 1; \
    fi;

vet:
    go vet $(VETPACKAGES)

# Cross compilation
build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v
build-win32:
        CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_386) -v
build-win64:
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_AMD64) -v

# docker-build:
#     echo "build the $(BINARY_LINUX) in docker"
#     docker run --rm -it -v "$(GOPATH)":/data/boyosoft/goBillProcess -w /data/boyosoft/goBillProcess/src/github.com/boyosoft/case_makefile devuser/gopher go build -o "$(BINARY_LINUX)" -v