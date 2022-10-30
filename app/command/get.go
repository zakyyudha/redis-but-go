package command

import (
	"fmt"
	"github.com/zakyyudha/redis-but-go/app/protocol"
	"github.com/zakyyudha/redis-but-go/app/storage"
	"net"
)

func Get(conn net.Conn, args []protocol.Value, storage storage.Storage) {
	if len(args) == 0 {
		conn.Write([]byte(fmt.Sprintf("-ERR cannot get empty key\r\n")))
		return
	}
	value, found := storage.Get(args[0].String())
	if found {
		conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)))
	} else {
		conn.Write([]byte("$-1\r\n"))
	}
}
