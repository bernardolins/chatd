package channel

import (
	"testing"
)

func TestNew(t *testing.T) {
	expectName := "General"
	c := New("General")
	gotName := c.Name()

	if expectName != gotName {
		t.Fatalf("Channel Constructor failed: expect %s, got %s", expectName, gotName)
	}
}

func TestAddUser(t *testing.T) {

}

func TestUserEvent(t *testing.T) {
	c := New("general")
	e := c.UserEvent("Fausto", "join", "general")
	expectUser := "Fausto"
	expectEvent := "join"
	expectValue := "general"
	expectChannel := "general"

	if expectChannel != e.Channel || expectUser != e.User || expectValue != e.Value || expectEvent != e.Action {
		t.Fatalf("UserEvent Failed: expect %s, %s, %s, %s, but got %s, %s, %s, %s", expectEvent, expectChannel, expectUser, expectValue, e.Action, e.Channel, e.User, e.Value)
	}
}
