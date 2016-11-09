package handler


import (
	"fmt"
	"strings"
)

func HandleKill(conn Conn) {
	conn.Close()
}

func HandleHelo(message string, conn Conn) {
	// conn.LocalAddr() returns "10.62.0.117:8000" for example
	// serverInfo[0] = 10.62.0.117
	// serverInfo[1] = 8000
	serverInfo := strings.Split(conn.localAddr(), ":")
	conn.Write([]byte(fmt.Sprintf("%sIP:%s\nPort:%s\nStudentID:%s\n", message, serverInfo[0], serverInfo[1], "13321218")))
}

func HandleOther(conn Conn) {
	// stub methond
	fmt.Println(conn)
}
