REL = r0
GIT_REVLIST = $(shell test -d .git && git rev-list --count HEAD)
GIT_DESCRIBE = $(shell test -d .git && git describe --always)
GIT_REF = $(shell test -d .git && git rev-parse --abbrev-ref HEAD)

ifneq "$(GIT_DESCRIBE)" ""
REL = r$(GIT_REVLIST).$(GIT_DESCRIBE)
endif

ifndef VERSION
ifneq "$(GIT_REF)" "master"
VERSION = $(REL)-$(GIT_REF)
else
VERSION = $(REL)
endif
endif

ifndef BUILDFLAGS
BUILDFLAGS = -i -v
endif

default: archivegen

all: test archivegen

archivegen:
	go build -ldflags "-X main.buildversion=$(VERSION)" $(BUILDFLAGS) bitbucket.org/multimfi/archivegen

install: archivegen
	install archivegen $(HOME)/.local/bin/archivegen

test:
	CGO_ENABLED=1 go test -race ./...

clean:
	rm -v archivegen

.PHONY: archivegen test
