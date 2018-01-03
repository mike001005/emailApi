#!/bin/bash

set -e
echo "" > coverage.txt


go get -u github.com/golang/lint/golint
go get github.com/GeertJohan/fgt
go get golang.org/x/tools/cmd/goimports

fgt goimports -l ./pkg/*
fgt golint ./pkg/...

#fgt goimports -l ./cmd/*
#fgt golint ./cmd/...

#fgt goimports -l ./internal/*
#fgt golint ./internal/...

for d in $(go list ./pkg/... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

