package arrays

import (
	"reflect"
	"testing"
)

func TestArrays(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d, but want %d, input %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
