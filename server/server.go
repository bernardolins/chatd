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
	server.ip = ip
	server.port = port

	fmt.Println("Starting server on %s:%d", server.ip, server.port)
	ln, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	server.listener = ln
}

func (server *Server) Stop() {
	server.listener.Close()
}

func (server *Server) Accept() {
	conn, err := server.listener.Accept()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go HandleConnection(conn)
}

func HandleConnection(conn net.Conn) {
	//handle new connections
}
