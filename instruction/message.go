package instruction

import (
	"fmt"
)

type Message struct{}

func (m Message) Run(target string) {
	fmt.Println(target)
}
