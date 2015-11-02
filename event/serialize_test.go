package event

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	b := []byte(`{"User":"Bernardo","Action":"Hello","Value":"World","Channel":"general"}`)

	e := Serialize(b)

	if e.User != "Bernardo" || e.Action != "Hello" || e.Value != "World" || e.Channel != "general" {
		t.Fatalf("Expect Bernardo, Hello, World, general, but got %s, %s, %s, %s", e.User, e.Action, e.Channel, e.Value)
	}
}
