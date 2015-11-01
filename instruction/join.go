package instruction

import (
	"fmt"
)

type Join struct{}

func (j Join) Run(target string) {
	fmt.Println(target)
}
