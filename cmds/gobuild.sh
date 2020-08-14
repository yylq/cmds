#!/usr/bin/bash
[[ -z $GOPATH ]] && GOPATH=$(pwd|cut -d '/' -f1-5)
export GOPATH
BIN="${GOPATH}bin"
export BIN
echo $GOPATH
echo $BIN
go build $1

