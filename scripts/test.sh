#!/bin/bash

set -e

HERE=$(dirname $0)
cd $HERE/..

mkdir -p reports

PKG_LIST=$(go list ./... | grep -v /vendor/)

# unit tests
go test $PKG_LIST
# memory sanitisation
#go test -short -msan $PKG_LIST
# race detection
go test -short -race $PKG_LIST

# TODO
# golint -set_exit_status ${PKG_LIST}

# coverage
for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "reports/${package##*/}.cov" "$package" ;
    go tool cover -html="reports/${package##*/}.cov" -o reports/${package##*/}.html
done

# go get gitlab.com/tslocum/godoc-static
