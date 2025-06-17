.PHONY: build
build:
	go build -trimpath -ldflags="-s -w" -o build/yarm ./cmd/yarm

.PHONY: test
test:
	go test -v ./...

