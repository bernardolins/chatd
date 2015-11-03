package action

import (
	"github.com/bernardolins/chatd/user"
)

func CreateUser(name string) {
	user := user.New(name)
	channel := controller.ByName("general")
	channel.AddUser(user)
}
