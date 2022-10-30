package command

import (
	"fmt"
	"github.com/zakyyudha/redis-but-go/app/protocol"
	"github.com/zakyyudha/redis-but-go/app/storage"
	"net"
	"strconv"
	"time"
)

func Set(conn net.Conn, args []protocol.Value, storage storage.Storage) {
	if len(args) == 0 {
		conn.Write([]byte(fmt.Sprintf("-ERR empty key value\r\n")))
		return
	}
	if len(args) > 2 {
		expiryStr := args[3].String()
		switch args[2].String() {
		case "px": // Handling for milliseconds
			expiryInMilliseconds, err := strconv.Atoi(expiryStr)
			if err != nil {
				conn.Write([]byte(fmt.Sprintf("-ERR PX value (%s) is not an integer\r\n", expiryStr)))
				break
			}
			storage.SetWithExpiry(args[0].String(), args[1].String(), time.Duration(expiryInMilliseconds)*time.Millisecond)
		case "ex": // Handling for seconds
			expiryInMilliseconds, err := strconv.Atoi(expiryStr)
			if err != nil {
				conn.Write([]byte(fmt.Sprintf("-ERR PX value (%s) is not an integer\r\n", expiryStr)))
				break
			}
			storage.SetWithExpiry(args[0].String(), args[1].String(), time.Duration(expiryInMilliseconds)*time.Second)
		default: // unknown option for set with expiry
			conn.Write([]byte(fmt.Sprintf("-ERR unknown option for set: %s\r\n", args[2].String())))
		}
	} else {
		if len(args) == 1 {
			conn.Write([]byte(fmt.Sprintf("-ERR cannot set empty value\r\n")))
			return
		}
		storage.Set(args[0].String(), args[1].String())
	}
	conn.Write([]byte("+OK\r\n"))
}
