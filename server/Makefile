GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
GOLANGCI=golangci-lint

run: ## go run main.go
	GO_ENV=develop $(GORUN) ./main.go

gtest: ## go test
	$(GOTEST) -v ./...

benchmark: ## go benchmark
	$(GOTEST) -bench -v ./...

doc: ## godoc http:6060
	$(GODOC) -http=:6060

build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) ./main.go

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'