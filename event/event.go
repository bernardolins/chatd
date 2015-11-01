// Defines Event type and handles event communication between modules
package event

type E struct {
	User    string
	Action  string
	Value   string
	Channel string
}

func New(user string, action string, value string, channel string) *E {
	e := new(E)
	e.User = user
	e.Action = action
	e.Value = value
	e.Channel = channel

	return e
}
