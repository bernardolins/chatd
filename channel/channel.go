package channel

import (
	"github.com/bernardolins/chatd/event"
	"github.com/bernardolins/chatd/peer"
	"github.com/chuckpreslar/emission"
)

type Channel struct {
	name         string
	eventHandler *emission.Emitter
}

func New(name string) *Channel {
	c := new(Channel)
	c.name = name
	c.eventHandler = emission.NewEmitter()

	return c
}

func (c *Channel) Name() string {
	return c.name
}

func (c *Channel) EventHandler() *emission.Emitter {
	return c.eventHandler
}

func (c *Channel) RegisterPeer(p *peer.Peer) {
	e := event.New(p.Identification(), "join", c.name, c.name)
	p.NotifyChannel(e, c.eventHandler)

	//Also binds the new peer to emitter, so it can listen to other peers' events
	p.BindChannel("join", c.eventHandler)
	p.BindChannel("message", c.eventHandler)
}

//func (c *Channel) NewEventFrom(u *peer.Peer, e *event.E) {
//	u.Notify(e, c.eventHandler)
//}

// Instantiates peer events on this channel
//func (c *Channel) PeerEvent(nick string, eventName string, value string) *event.E {
//	return event.New(nick, eventName, value, c.name)
//}
