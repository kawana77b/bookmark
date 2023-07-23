VERSION := v0.0.7

build:
	@go build -o ./dist/

build-r:
	@goreleaser release --snapshot --rm-dist

install:
	@go install -ldflags="-s -w -extldflags=-static -X main.Version=$(VERSION)"

release:
	@git checkout main
	@git tag -a $(VERSION) -m "$(VERSION)"
	@git push --tags origin

credits:
	@gocredits . > CREDITS

test:
	@go test -v ./...

.PHONY: build build-r install release credits test
