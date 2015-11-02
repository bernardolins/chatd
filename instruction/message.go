package instruction

import (
	"github.com/bernardolins/chatd/event"
)

type Message struct{}

func (m Message) Run(target string, event *event.E) {
}
