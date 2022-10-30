package main

import (
	"bufio"
	"fmt"
	cmd "github.com/zakyyudha/redis-but-go/app/command"
	"github.com/zakyyudha/redis-but-go/app/protocol"
	"github.com/zakyyudha/redis-but-go/app/storage"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379: ", err.Error())
		os.Exit(1)
	}

	// Init storage
	memoryStorage := storage.NewMemoryStorage()

	fmt.Println("Server started on 0.0.0.0:6379")
	for true {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// Handle concurrent connection
		go handleConnection(conn, memoryStorage)
	}
}

func handleConnection(conn net.Conn, storage storage.Storage) {
	defer conn.Close()
	for {
		value, err := protocol.NewRESPProtocol(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				conn.Close()
				fmt.Println("Client disconnected from server!")
				return
			}
			fmt.Println("Error decoding protocol: ", err.Error())
		}

		command := value.Array()[0].String()
		args := value.Array()[1:]

		switch command {
		case "ping":
			cmd.Ping(conn)
			break
		case "echo":
			cmd.Echo(conn, args)
			break
		case "set":
			cmd.Set(conn, args, storage)
			break
		case "get":
			cmd.Get(conn, args, storage)
			break
		default:
			conn.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
			break
		}
	}
}
