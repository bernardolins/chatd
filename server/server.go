package server

import (
	"fmt"
	"github.com/bernardolins/chatd/channel"
	"github.com/bernardolins/chatd/event"
	"github.com/bernardolins/chatd/user"
	"net"
	"os"
)

type Server struct {
	ip                string
	port              string
	listener          net.Listener
	channelController *channel.Controller
}

func New(ip string, port string) *Server {
	s := new(Server)
	s.ip = ip
	s.port = port

	s.channelController = channel.NewController()
	s.channelController.NewChannel("original")

	return s
}

func (server *Server) Up() {
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

	userBuff := make([]byte, 1024)

	// Wait informations about new user
	conn.Read(userBuff)
	e := event.Serialize(userBuff)

	// While client does not send the info the server should not create the user
	var u *user.User
	if e.Action == "createUser" {
		//instruction.Controller.InstructionByName["createUser"].Run(e.User, e)
		fmt.Println("Creating user %s", e.User)
		u = user.New(e.Value)
		c := server.channelController.SelectChannel("general")
		c.AddUser(u)
	}

	server.HandleConnection(conn, u)
}

// Handles User Connection
func (server *Server) HandleConnection(conn net.Conn, user *user.User) {
	go server.HandleIncomingRequest(conn, user)
}

func (server *Server) HandleIncomingRequest(conn net.Conn, user *user.User) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err.Error())
	}

	e := event.Serialize(buffer)
	channel := server.channelController.SelectChannel(e.Channel)
	channel.NewEventFrom(user, e)
}
