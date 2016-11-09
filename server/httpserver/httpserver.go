package httpserver

import (
	"net"
	"os"
	"github.com/TomasBarry/Lab2/server/handler"
)

const (
	CONN_TYPE = "tcp"
	HELO_COMMAND = "HELO "
	KILL_COMMAND = "KILL_SERVICE\n"
)

func CreateServer(port string) {
	listener, _ := net.Listen(CONN_TYPE, ":" + port)
	// wait for connections
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func readSocket(conn Conn) message {
	buffer := make([]byte, 1024)
	readLength, _ := conn.Read(buff)
	message := string(buff[:readLength])
}

func handleConnection(conn Conn) {
	// persist the socket connection
	for {
		switch message := readSocket(conn) {
		case message.HasPrefix(HELO_COMMAND):
			handler.HandleHelo(message, conn)
		case message == KILL_COMMAND:
			handler.HandleKill(conn)
		default:
			handler.HandleOther(conn)
		}

	}
}
