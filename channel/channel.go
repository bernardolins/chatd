package channel

import (
	"github.com/bernardolins/chatd/event"
	"github.com/bernardolins/chatd/user"
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

func (c *Channel) AddUser(u *user.User) {
	//New user notifies everybody  binded to this emitter he is on the channel
	e := c.UserEvent(u.Nickname(), "join", c.name)
	u.Notify(e, c.eventHandler)

	//Also binds the new user to emitter, so it can listen to other users' events
	u.Bind("join", c.eventHandler)
	u.Bind("message", c.eventHandler)
}

func (c *Channel) NewEventFrom(u *user.User, e *event.E) {
	u.Notify(e, c.eventHandler)
}

// Instantiates user events on this channel
func (c *Channel) UserEvent(nick string, eventName string, value string) *event.E {
	return event.New(nick, eventName, value, c.name)
}
