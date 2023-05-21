package structs

import (
	"testing"
)

func TestPerimiter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimiter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestAerea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g, but want %g", got, want)
		}
	}

	t.Run("run rectangle test", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, rect, want)
	})

	t.Run("run circle test", func(t *testing.T) {
		circle := Circle{10}

		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}
