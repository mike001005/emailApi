#!/bin/bash

set -e
echo "" > coverage.txt

go get golang.org/x/tools/cmd/goimports
go get -u github.com/golang/lint/golint
go get github.com/GeertJohan/fgt
go get golang.org/x/tools/cmd/goimports
go get -v -u github.com/gorilla/mux
go get -v -u github.com/sendgrid/sendgrid-go
go get -v -u gopkg.in/mailgun/mailgun-go.v1


fgt goimports -l ./pkg/*
fgt golint ./pkg/...

for d in $(go list ./pkg/... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

curl -s https://codecov.io/bash | bash -s - -t 4ffaaabc-09b8-4ca5-9c22-6eff395fc643

