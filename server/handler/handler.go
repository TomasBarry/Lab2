package handler


import (
	"fmt"
	"net"
	"strings"
)

func HandleKill(conn net.Conn) {
	conn.Close()
}

func HandleHelo(message string, conn net.Conn) {
	// conn.LocalAddr() returns "10.62.0.117:8000" for example
	// serverInfo[0] = 10.62.0.117
	// serverInfo[1] = 8000
	serverInfo := strings.Split(conn.LocalAddr().String(), ":")
	conn.Write([]byte(fmt.Sprintf("%sIP:%s\nPort:%s\nStudentID:%s\n", message, serverInfo[0], serverInfo[1], "13321218")))
}

func HandleOther(conn net.Conn) {
	// stub methond
	fmt.Println(conn)
}
