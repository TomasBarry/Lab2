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
	IP_ADDR = "10.62.0.117"
	STUDENT_NUMBER = "13321218"
)

var (
	activeThreads = 0
)

func handleKillCommand(conn net.Conn) {
	conn.Close()
	activeThreads -= 1
}

func handleHeloCommand(message string, port string, conn net.Conn) {
	// "HELO text\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"
	var response bytes.Buffer
	response.WriteString(message)
	response.WriteString(fmt.Sprintf("IP:%s\n", IP_ADDR))
	response.WriteString(fmt.Sprintf("Port:%s\n", port))
	response.WriteString(fmt.Sprint("StudentID:%s\n", STUDENT_NUMBER))
	conn.Write([]byte(response.String()))
}

func handleOtherCommand() {
	// Stub function
}

func handleConnection(conn net.Conn) {
	validHello := regexp.MustCompile(`HELO *`)
	validKill := regexp.MustCompile(`KILL_SERVICE *`)
	port := conn.LocalAddr().String()[12:]
	// wait for messages on the socket
	going := true
	for going {
		fmt.Println("Addr: ", conn.LocalAddr())
		buff := make([]byte, 1024)
		readLen, e := conn.Read(buff)
		handleError(e)
		message := string(buff[:readLen])
		if validKill.MatchString(message) {
			going = false
			handleKillCommand(conn)
		} else if  validHello.MatchString(message) {
			handleHeloCommand(message, port, conn)
		} else {
			handleOtherCommand()
		}
	}
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
	listener, e := net.Listen(CONN_TYPE, ":" + args[0])
	handleError(e)
	// wait for new clients to connect
	for {
		conn, e := listener.Accept()
		fmt.Println("A new connection")
		handleError(e)
		if activeThreads < MAX_THREAD_POOL {
			activeThreads += 1
			go handleConnection(conn)
		}
	}
}

