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

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.hasArea

		if got != want {
			t.Errorf("%#vgot %g, but want %g", tt.shape, got, want)
		}
	}
}
