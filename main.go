package main

import (
	"fmt"
	"github.com/bernardolins/chatd/server"
)

func main() {
	server := server.New("0.0.0.0", "9090")
	server.Up()

	for {
		server.Accept()
	}

	var input string
	fmt.Scanln(&input)

}
