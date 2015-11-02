package server

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	ip       string
	port     byte
	listener net.Listener
}

func (server *Server) Up(ip string, port byte) {
	ln, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server.listener = ln
}
