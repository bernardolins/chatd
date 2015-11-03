package server

import (
	"bufio"
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
	s.channelController.NewChannel("general")

	return s
}

func (server *Server) Up() {
	fmt.Println("Starting server on ", server.ip, ":", server.port)
	ln, err := net.Listen("tcp", server.ip+":"+server.port)

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
	fmt.Println("Accept")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go server.HandleConnection(conn)
}

// Handles User Connection
func (server *Server) HandleConnection(conn net.Conn) {
	for {
		userBuff, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println(userBuff)
		e, err := event.Serialize([]byte(userBuff))

		if err != nil {
			conn.Close()
			return
		}

		conn.Write([]byte(userBuff))

		var u *user.User
		if e.Action == "createUser" {
			u = user.New(e.Value)
			c := server.channelController.SelectChannel("general")
			c.AddUser(u)
		} else {
			channel := server.channelController.SelectChannel(e.Channel)
			channel.NewEventFrom(u, e)
		}

		UserResponse(conn, u)
	}
}

func UserResponse(conn net.Conn, u *user.User) {
	fmt.Println(u.Nickname(), len(u.EventList()))
	for _, e := range u.EventList() {
		de := event.Deserialize(e)
		fmt.Println(string(de))
		conn.Write(de)
	}
}
