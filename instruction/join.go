package instruction

import (
	"github.com/bernardolins/chatd/event"
)

type Join struct{}

func (j Join) Run(target string, event *event.E) {
}
