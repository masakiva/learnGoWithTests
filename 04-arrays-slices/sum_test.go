package main

import (
	"reflect"
	"testing"
)

// why it's a good idea to make a copy of a slice after slicing a very large slice: https://play.golang.org/p/Poth8JS28sc
func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} // Go blog post on slices: https://blog.golang.org/go-slices-usage-and-internals

		got := Sum(numbers)
		want := 15

		assertCorrectResult(t, got, want, numbers)
	})
}

func assertCorrectResult(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
