package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("r", 3)
	want := "rrr"

	if got != want {
		t.Errorf("got %q as we want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
