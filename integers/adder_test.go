package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("expected %d but got %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(5, 1)
	fmt.Println(sum)
	// Output: 6
}
