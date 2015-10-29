package user

import (
	"github.com/chuckpreslar/emission"
)

type User struct {
	nickname string
	ip       string
	listener *emission.Emitter
}

//Instantiates a new User
func New(nickname string, ip string) *User {
	u := new(User)
	u.nickname = nickname
	u.ip = ip
	u.listener = emission.NewEmitter()

	return u
}

func (u *User) Listener() *emission.Emitter {
	return u.listener
}

func (u *User) Ip() string {
	return u.ip
}

func (u *User) Nickname() string {
	return u.nickname
}
