.PHONY: info tools test

GOTOOLS = \
	github.com/tcnksm/ghr \
	github.com/mitchellh/gox

PACKAGES=$(shell go list ./... | grep -v /vendor/)

GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --tags | grep -oP "v[0-9]+(\.[0-9]+)*")
VERSION=$(shell git describe --tags | grep -oP "v[0-9]+(\.[0-9]+)*" | sed 's/v//')

info:
	@echo "Settings:"
	@echo "GIT_COMMIT: $(GIT_COMMIT)"
	@echo "GIT_TAG:    $(GIT_TAG)"
	@echo "VERSION:    $(VERSION)"

tools:
	go get -u -v $(GOTOOLS)

test:
	go test -v $(PACKAGES) -cover


