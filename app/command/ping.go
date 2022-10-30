package command

import "net"

func Ping(conn net.Conn) {
	conn.Write([]byte("+PONG\r\n"))
}
