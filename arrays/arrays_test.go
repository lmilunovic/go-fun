package arrays

import "testing"

func TestArrays(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("want %d got %d, given %v", got, want, numbers)
	}
}
