package event

import (
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	event := New("Bernardo", "Hello", "World", "general")
	b := Deserialize(event)
	expectByte := []byte(`{"User":"Bernardo","Action":"Hello","Value":"World","Channel":"general"}`)
	expect := reflect.DeepEqual(b, expectByte)

	if !expect {
		t.Fatalf("Event serialization failed! Expected %q, but got %q", expectByte, b)
	}
}
