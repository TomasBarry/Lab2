package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	CONN_TYPE = "tcp"
	MAX_THREAD_POOL = 1
)

var (
	activeThreads = 0
)

func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection ", activeThreads)
	buff, _ := ioutil.ReadAll(conn)
	fmt.Println(string(buff))
	activeThreads -= 1
	fmt.Println("Handled connection ", activeThreads)
}

func main() {
	// accept command line arguements where args[0] is the port number to run on
	args := os.Args[1:]
	listener, _ := net.Listen(CONN_TYPE, ":" + args[0])
	// wait for new clients to connect
	for {
		conn, _ := listener.Accept()
		if activeThreads < MAX_THREAD_POOL {
			fmt.Println("Less than max number of threads")
			activeThreads += 1
			go handleConnection(conn)
		} else {
			fmt.Println("Too many threads already")
		}
	}
}

