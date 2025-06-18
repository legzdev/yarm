MODULE=$(shell cat go.mod | grep module | head -n 1 | cut -d " " -f 2)

VERSION=$(shell git tag --sort=-creatordate | head -n 1)
ifeq ($(VERSION), $(echo))
	VERSION=unknown
endif

.PHONY: build
build:
	go build -trimpath -ldflags="-s -w -X $(MODULE).Version=$(VERSION)" -o build/yarm ./cmd/yarm

.PHONY: test
test:
	go test -v ./...

