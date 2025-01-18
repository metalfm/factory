lint:
	go run -modfile=tools/go.mod github.com/golangci/golangci-lint/cmd/golangci-lint run --allow-parallel-runners
.PHONY: lint

gen:
	go generate -v ./...
.PHONY: gen

test:
	go test -v -race ./...
.PHONY: test

