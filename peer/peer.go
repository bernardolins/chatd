package peer

import (
	"github.com/bernardolins/chatd/event"
	"github.com/chuckpreslar/emission"
	"net"
)

type Peer struct {
	identification string
	connection     net.Conn
}

func New(identification string, connection net.Conn) *Peer {
	p := new(Peer)
	p.connection = connection
	p.identification = identification

	return p
}

// Getter for Connection
func (p *Peer) Connection() net.Conn {
	return p.connection
}

func (p *Peer) Identification() string {
	return p.identification
}

// Binds Peer to an event, using some external event emitter
func (p *Peer) BindChannel(e string, emitter *emission.Emitter) {
	emitter.On(e, p.receiveNotification)
}

// Trigger events on an external event emitter
func (u *Peer) NotifyChannel(e *event.E, emitter *emission.Emitter) {
	emitter.Emit(e.Action, e)
}

func (p *Peer) receiveNotification(e *event.E) {
	byteEvent := event.Deserialize(e)
	p.connection.Write(byteEvent)
}
