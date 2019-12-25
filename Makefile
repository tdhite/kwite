# Makefile
#
# Copyright (c) 2019-2020 VMware, Inc.
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

GOFILES = $(shell find . -type f -not -path '.venddor/*' -name '*.go')

default: cmd/kwite/kwite

container: cmd/kwite/kwite
	cd build/image; ./build.sh
.PHONY: container

cmd/kwite/kwite: go.mod vet fmt $(GOFILES)
	cd cmd/kwite; GOARCH=amd64 CGO_ENABLED=0 go build -a --installsuffix cgo kwite.go

test: go.mod
	go test ./...
.PHONY: test

fmt:
	go fmt ./...

vet: go.mod
	go vet ./...

go.mod:
	go mod init github.com/tdhite/kwite

clean:
	cd cmd/kwite && \
	go clean && \
	rm -f kwite && \
	go clean -modcache
.PHONY: clean
