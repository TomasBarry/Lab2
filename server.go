package main

import (
    "fmt"
    "net"
)

const (
    MaxNumberOfThreads = 3
)

var (
    activeThreads = 0
)

func handleConnection(c net.Conn) {
    fmt.Println("Handling connection ", activeThreads)
    activeThreads -= 1
    fmt.Println("Handled connection ", activeThreads)
}

func main() {
    ln, _ := net.Listen("tcp", ":8000")
    for {
        conn, _ := ln.Accept()
        if activeThreads < MaxNumberOfThreads {
            fmt.Println("Less than max number of threads")
            activeThreads += 1 
            go handleConnection(conn)
        } else {
            fmt.Println("Too many threads already")
        }
    }
}
