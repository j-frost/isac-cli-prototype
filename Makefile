build:
	go build -o isac-cli-prototype cmd/isac-cli-prototype/main.go

install:
	go build -o $(GOBIN)/isac-cli-prototype cmd/isac-cli-prototype/main.go

run:
	go run cmd/isac-cli-prototype/main.go

goimports:
	go get golang.org/x/tools/cmd/goimports

format: goimports
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +
