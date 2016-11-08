#!/bin/bash

# Install Go on the current machine
export PATH=$PATH:Dependencies/go/bin
export GOPATH=Dependencies/gowork
export PATH=$PATH:$GOPATH/bin
go get github.com/tools/godep
apt-get install bzr -y

go install Server/server.go
