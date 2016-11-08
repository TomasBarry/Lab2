package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"bytes"
)

const (
	CONN_TYPE = "tcp"
	MAX_THREAD_POOL = 1
	KILL_COMMAND = "KILL_SERVICE\n"
	HELLO_COMMAND ="HELO text\n"
)

var (
	activeThreads = 0
)

func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection ", activeThreads)
	buff, e := ioutil.ReadAll(conn)
	handleError(e)
	message := string(buff)
	switch message {
	case KILL_COMMAND:
		os.Exit(0)
	case HELLO_COMMAND:
		// "HELO text\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"
		var response bytes.Buffer
		response.WriteString(message)
		response.WriteString("IP:10.62.0.117\n")
		response.WriteString("Port:8000\n")
		response.WriteString("StudentID:13321218\n")
		conn.Write([]byte(response.String()))
		fmt.Println(response.String())
	default:
		// This would handle all other messages
		fmt.Println(message)
	}
	activeThreads -= 1
	fmt.Println("Handled connection ", activeThreads)
}


func handleError(e error) {
	// if an error exits, exit the program
	if e != nil {
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

