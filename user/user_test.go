package user

import (
	"github.com/chuckpreslar/emission"
	"testing"
)

func TestNew(t *testing.T) {
	expectIp := "1.2.3.4"
	expectNick := "Fausto"

	u := New("Fausto", "1.2.3.4")
	gotIp := u.Ip()
	gotNick := u.Nickname()

	if expectNick != gotNick || expectIp != gotIp {
		t.Fatalf("Unnable to construct new user: expect %s and %s, got %s, %s", expectIp, expectNick, gotIp)
	}
}
