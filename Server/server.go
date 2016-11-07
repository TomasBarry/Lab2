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
	buff, e := ioutil.ReadAll(conn)
	handleError(e)
	message := string(buff)
	activeThreads -= 1
	fmt.Println("Handled connection ", activeThreads)
}

func handleError(e error) {
	// if an error exits, exit the program
	if e {
		fmt.Println("An error occurred: %s", e.Error())
		os.Exit(1)
	}
}

func main() {
	// accept command line arguements where args[0] is the port number to run on
	args := os.Args[1:]
	// create the listener
	listener, e := net.Listen(CONN_TYPE, ":" + args[0])
	handleError(e)
	// wait for new clients to connect
	for {
		conn, e := listener.Accept()
		handleError(e)
		if activeThreads < MAX_THREAD_POOL {
			fmt.Println("Less than max number of threads")
			activeThreads += 1
			go handleConnection(conn)
		} else {
			fmt.Println("Too many threads already")
		}
	}
}

