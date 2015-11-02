package event

import (
	"encoding/json"
)

func Serialize(b []byte) *E {
	var event E
	json.Unmarshal(b, &event)

	return &event
}
