#!/bin/bash

# Install Go on the current machine
export PATH=$PATH:$PWD/Dependencies/gowork/bin
export GOPATH=$PWD/Dependencies/gowork/src
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin
go get github.com/tools/godep
apt-get install bzr -y

go get github.com/TomasBarry/Lab2/Server
