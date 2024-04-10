export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

.PHONY: build
build:
	go build -o app .
