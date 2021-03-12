PROJECT_NAME := "demo"
MAIN_FILE_PAHT := "main.go"
PKG := "github.com/infraboard/demo"
IMAGE_PREFIX := "github.com/infraboard/demo"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v redis)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}
	
test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go fmt ./...
	@sh ./script/build.sh local dist/${PROJECT_NAME} ${MAIN_FILE_PAHT} ${IMAGE_PREFIX} ${PKG}

linux: ## Linux build
	@sh ./script/build.sh linux dist/${PROJECT_NAME} ${MAIN_FILE_PAHT} ${IMAGE_PREFIX} ${PKG}
	
run: install codegen dep build ## Run Server
	@./dist/${PROJECT_NAME} start

clean: ## Remove previous build
	@go clean .
	@rm -f dist/${PROJECT_NAME}

install: ## Install depence go package
	@go install github.com/golang/protobuf/protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@go install github.com/infraboard/protoc-gen-go-ext
	@go install github.com/infraboard/mcube/cmd/protoc-gen-go-http

codegen: ## Init Service
	@protoc -I=.  -I${GOPATH}/src --go-ext_out=. --go-ext_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} --go-http_out=. --go-http_opt=module=${PKG} pkg/*/pb/*.proto
	@go generate ./...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'