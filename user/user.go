package user

import (
	"github.com/bernardolins/chatd/event"
	"github.com/chuckpreslar/emission"
)

type User struct {
	nickname string
	ip       string
}

func New(nickname string, ip string) *User {
	u := new(User)
	u.nickname = nickname
	u.ip = ip

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

// Binds User to an event, using some external event emitter
func (u *User) Bind(e string, emitter *emission.Emitter) {
	emitter.On(e, event.Handle)
}

// Trigger events on an external event emitter
func (u *User) Notify(e *event.E, emitter *emission.Emitter) {
	emitter.Emit(e.Action, e)
}
