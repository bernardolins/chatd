package user

import (
	"github.com/chuckpreslar/emission"
)

type User struct {
	nickname     string
	ip           string
	eventHandler *emission.Emitter
}

//Instantiates a new User
func New(nickname string, ip string) *User {
	u := new(User)
	u.nickname = nickname
	u.ip = ip
	u.eventHandler = emission.NewEmitter()

	return u
}

func (u *User) EventHandler() *emission.Emitter {
	return u.eventHandler
}

func (u *User) Ip() string {
	return u.ip
}

func (u *User) Nickname() string {
	return u.nickname
}
