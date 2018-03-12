GIT_REF := $(shell git rev-parse --short HEAD || echo unsupported)
BUILDFLAGS := -ldflags "-X main.gitRef=$(GIT_REF)"

build:
	go build -v $(BUILDFLAGS)

release:
	goreleaser --rm-dist

ansible:
	ansible-playbook -i deploy/ansible/hosts deploy/ansible/deploy.yml
