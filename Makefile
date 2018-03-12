GIT_REF := $(shell git rev-parse --short HEAD || echo unsupported)
BUILDFLAGS := -ldflags "-X main.gitRef=$(GIT_REF)"

build:
	go build -v $(BUILDFLAGS)

release:
	goreleaser --rm-dist
