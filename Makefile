REL = r0

GIT_DESCRIBE := $(shell test -d .git && git describe --always)
GIT_REVLIST := $(shell test -d .git && git rev-list --count HEAD)
GIT_DIRTY := $(shell test -d .git && git diff-index --quiet HEAD || date '+-dirty-%s')
GIT_REF := $(shell test -d .git && git rev-parse --abbrev-ref HEAD)

ifneq "$(GIT_DESCRIBE)" ""
REL := r$(GIT_REVLIST).$(GIT_DESCRIBE)
endif

ifndef VERSION
ifneq "$(GIT_REF)" "master"
VERSION := $(REL)-$(GIT_REF)$(GIT_DIRTY)
else
VERSION := $(REL)$(GIT_DIRTY)
endif
endif

ifndef BUILDFLAGS
BUILDFLAGS = -i -v
endif

default: archivegen

all: test archivegen

archivegen:
	@go build -ldflags "-X main.buildversion=$(VERSION)" $(BUILDFLAGS) github.com/multimfi/archivegen

goinstall:
	go install -ldflags "-X main.buildversion=$(VERSION)" -v github.com/multimfi/archivegen

install: archivegen
	install archivegen $(HOME)/.local/bin/archivegen

test:
	@go test $(shell go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./...|grep -v /vendor/)

clean:
	rm -v archivegen

.PHONY: archivegen
