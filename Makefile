#! /usr/bin/make
#
# Makefile for gorma
#
# Targets:
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./... | grep -v /example)
DEPEND=\
	github.com/volatiletech/inflect \
	github.com/seesaa/goa/... \
	github.com/seesaa/goa.design/tools/mdc \
	github.com/golang/lint/golint \
	github.com/onsi/ginkgo \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega \
	github.com/jinzhu/inflection \
	github.com/kr/pretty \
	golang.org/x/tools/cmd/goimports

.PHONY: goagen

all: depend lint test

docs:
	@go get -v github.com/spf13/hugo
	@cd /tmp && git clone https://github.com/seesaa/goa.design && cd goa.design && rm -rf content/reference public && make docs && hugo
	@rm -rf public
	@mv /tmp/goa.design/public public
	@rm -rf /tmp/goa.design

depend:
	@export GO111MODULE=on && go get -v github.com/seesaa/goa.design/tools/godoc2md
	@go get -v $(DEPEND)
	@go install $(DEPEND)


lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi


test:
	@ginkgo -r  --failOnPending  --race -skipPackage vendor
