#!/bin/bash

# Download Go
wget https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.6.3.linux-amd64.tar.gz


# Set up environment variables for Go
mkdir ~/gowork

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/gowork
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin


# Install Go Dependencies
go get github.com/tools/godep
apt-get install bzr -y


# Get the repo
go get github.com/TomasBarry/Lab2


# Install the repo
go install lab2.go
