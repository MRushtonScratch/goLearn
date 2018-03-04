package methods

import
(
	"testing"
)


func TestCalcDistanceOfPath(t *testing.T) {
	var expected float64 = 15.142135623730951

	path := Path{
		Point{0,0},
		Point{5, 5},
		Point{10,10},
	}

	d := path.Distance()

	if d != expected {
		t.Errorf("PAth distance should be %d, was %d", expected, d)
	}
}