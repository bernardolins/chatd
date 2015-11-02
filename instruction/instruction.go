package instruction

import (
	"github.com/bernardolins/chatd/event"
)

type Instruction interface {
	Run(string, *event.E)
}

func Call(target string, event *event.E) {
	instruction := InstructionByName(event.Action)
	instruction.Run(target, event)
}
