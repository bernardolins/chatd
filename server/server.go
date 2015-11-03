package server

import (
	"bufio"
	"fmt"
	"github.com/bernardolins/chatd/channel"
	"github.com/bernardolins/chatd/event"
	"github.com/bernardolins/chatd/peer"
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
		return
	}

	go server.HandleConnection(conn)
}

func (server *Server) HandleConnection(conn net.Conn) {
	peer := new(peer.Peer)

	for {
		e, err := receiveAction(conn)
		if err != nil {
			fmt.Println(err.Error())
		}

		if e.Action == "connect" {
			peer = server.initializePeer(e.User, conn)
		}

		if peer != nil {
			handleConnectionError(conn, err)
			channel := server.channelController.SelectChannel(e.Channel)
			peer.NotifyChannel(e, channel.EventHandler())
		}
	}
}

func (server *Server) initializePeer(identification string, conn net.Conn) *peer.Peer {
	p := peer.New(identification, conn)
	channel := server.channelController.SelectChannel("general")
	channel.RegisterPeer(p)
	return p
}

func receiveAction(conn net.Conn) (*event.E, error) {
	userBuff, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(userBuff)

	if err != nil {
		return nil, err
	}

	e, serializeError := event.Serialize([]byte(userBuff))

	if serializeError != nil {
		fmt.Println(serializeError.Error())
		return nil, serializeError
	}

	_, err = conn.Write([]byte(userBuff))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return e, err
}

func handleConnectionError(conn net.Conn, err error) {
	if err != nil {
		conn.Close()
	}
}
