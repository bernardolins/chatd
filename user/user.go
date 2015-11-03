package user

import (
	"fmt"
	"github.com/bernardolins/chatd/event"
	"github.com/chuckpreslar/emission"
)

type User struct {
	nickname  string
	ip        string
	eventList []*event.E
}

func New(nickname string) *User {
	u := new(User)
	u.nickname = nickname

	u.eventList = make([]*event.E, 0)

	return u
}

// Getter for Ip
func (u *User) Ip() string {
	return u.ip
}

// Getter for Nickname
func (u *User) Nickname() string {
	return u.nickname
}

func (u *User) EventList() []*event.E {
	return u.eventList
}

// Binds User to an event, using some external event emitter
func (u *User) Bind(e string, emitter *emission.Emitter) {
	emitter.On(e, u.HandleEvent)
}

// Trigger events on an external event emitter
func (u *User) Notify(e *event.E, emitter *emission.Emitter) {
	emitter.Emit(e.Action, e)
}

func (u *User) HandleEvent(e *event.E) {
	fmt.Println(e.User, e.Action, e.Value, e.Channel)
	u.eventList = append(u.eventList, e)
}
