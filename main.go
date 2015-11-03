package main

import (
	"github.com/bernardolins/chatd/server"
)

func main() {
	server := server.New("127.0.0.1", "9090")
	server.Up()

	for {
		server.Accept()
	}
}
