#!/bin/bash

# Install Go on the current machine
export PATH=$PATH:$PWD/Dependencies/go/bin
export GOPATH=$PWD/Dependencies/gowork
export PATH=$PATH:$GOPATH/bin
go get github.com/tools/godep
apt-get install bzr -y

go install $PWD/Server/HttpServer/httpserver.go
