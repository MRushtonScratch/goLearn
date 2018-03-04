package methods

import
(
	"testing"
)


func TestCalcDistanceOfPath(t *testing.T) {
	var expected float64 = 19.62266869472574

	path := Path{
		Point{1,1},
		{5, 1},
		{10,1},
		{12,1},
	}

	d := path.Distance()

	if d != expected {
		t.Errorf("PAth distance should be %d, was %d", expected, d)
	}
}