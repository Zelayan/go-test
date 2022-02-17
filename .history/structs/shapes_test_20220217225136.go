package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle {
		10.0, 10.0,
	}

	got := Perimeter(rectangle)	
	want := 40.0
	if (want != got) {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    }

	t.Run("tes strangle", func(t *testing.T) {
		rectangle := Rectangle {
			10.0, 10.0,
		}
		got := rectangle.Area()
		want := 100.0
		checkArea(t, )
	})

	t.Run("test Circle", func(t *testing.T) {
		circle := Circle {
			10,
		}
		got2 := circle.Area()
		want2 := 314.1592653589793
		if (want2 != got2) {
			t.Errorf("got %f want %f", got2, want2)
		}
	})
	
}