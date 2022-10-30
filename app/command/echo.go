package command

import (
	"fmt"
	"github.com/zakyyudha/redis-but-go/app/protocol"
	"net"
)

func Echo(conn net.Conn, args []protocol.Value) {
	conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0].String()), args[0].String())))
}
