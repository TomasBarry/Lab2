package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"bytes"
)

const (
	CONN_TYPE = "tcp"
	MAX_THREAD_POOL = 10
)

var (
	activeThreads = 0
)

func handleConnection(conn net.Conn) {
	for {
		// fmt.Println("Addr: ", conn.LocalAddr())
		buff := make([]byte, 1024)
		readLen, e := conn.Read(buff)
		//fmt.Println("Message is : ", string(buff))
		handleError(e)
		message := string(buff[:readLen])
		//fmt.Println(message)
		validHello := regexp.MustCompile(`HELO *`)
		validKill := regexp.MustCompile(`KILL_SERVICE *`)
		if validKill.MatchString(message) {
			fmt.Println("Gonna kill the server")
			conn.Close()
			os.Exit(0)
		} else if  validHello.MatchString(message) {
			fmt.Println("Got a helo message")
			// "HELO text\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"
			var response bytes.Buffer
			response.WriteString(message)
			response.WriteString("IP:10.62.0.117\n")
			response.WriteString("Port:8000\n")
			response.WriteString("StudentID:13321218\n")
			conn.Write([]byte(response.String()))
			handleError(e)
			//fmt.Println(response.String())
		} else {
			fmt.Println("Got a random message")
			// This would handle all other messages
			//fmt.Println(message)
		}
	}
	//fmt.Println("Handled connection ", activeThreads)
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
		fmt.Println("A new connection")
		handleError(e)
		if activeThreads < MAX_THREAD_POOL {
			//fmt.Println("Less than max number of threads")
			activeThreads += 1
			go handleConnection(conn)
		} else {
			//fmt.Println("Too many threads already")
		}
	}
}

