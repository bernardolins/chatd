package main

import (
	"github.com/bernardolins/chatd/server"
)

func main() {
	server := server.New("0.0.0.0", "9090")
	server.Up()

	server.Accept()
}
