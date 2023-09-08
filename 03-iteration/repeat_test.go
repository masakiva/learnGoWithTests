package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("r")
	want := "rrrrr"

	if got != want {
		t.Errorf("got %q as we want %q", got, want)
	}
}
