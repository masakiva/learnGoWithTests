package main

import (
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if spySleeper.Calls != countdownStart {
		t.Errorf("not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.Calls)
	}
}
