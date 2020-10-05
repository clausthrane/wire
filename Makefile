PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./.mocks/*")

tools:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint

test:
	go test -race ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o=cover.html

fmt:
	go fmt ./...
	goimports -w $(FILES)

lint:
	golint $(PACKAGES)

vet:
	go vet ./...
