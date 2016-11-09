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

func createServer() {
	// accept command line arguements where args[0] is the port number to run on
	args := os.args[1:]
	listener, _ := net.Listen(CONN_TYPE, ":" + args[0])
	// wait for connections
	for {
		conn, _ listener.Accept()
		go handleConnection(conn)
	}
}

func readSocket(conn Conn) message string {
	buffer := make([]byte, 1024)
	readLength, _ := conn.Read(buff)
	message := string(buff[:readLength])
}

func handleConnection(conn Conn) {
	// persist the socket connection
	for {
		switch message := readSocket(conn) {
		case message.HasPrefix(HELO_COMMAND):
			handler.HandleHelo(conn)
		case message == KILL_COMMAND:
			handler.HandleKill(conn)
		default:
			handler.HandleOther(conn)
		}

	}
}
