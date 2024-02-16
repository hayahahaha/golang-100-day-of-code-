package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Testing rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f ", got, want)
		}
	})
	t.Run("Testing circles", func(t *testing.T) {

	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f ", got, want)
		}
	}
	t.Run("Testing rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 3.0}
		checkArea(t, rectangle, 36.0)
	})
	t.Run("Testing circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
