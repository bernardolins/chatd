package event

import (
	"encoding/json"
	"fmt"
)

func Deserialize(event *E) []byte {
	b, err := json.Marshal(event)

	if err != nil {
		fmt.Println("Warning: Could not deserialize event")
		fmt.Println(err.Error())
		return []byte(`{}`)
	}

	return b
}
