package main

import (
	"github.com/TomasBarry/Lab2/server/httpserver"
	"os"
)

func main() {
	httpserver.CreateServer(os.Args[1])
}
