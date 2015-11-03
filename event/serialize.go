package event

import (
	"encoding/json"
	"fmt"
)

func Serialize(b []byte) (*E, error) {
	var event E
	err := json.Unmarshal(b, &event)

	if err != nil {
		fmt.Println(err.Error())
	}
	return &event, err
}
