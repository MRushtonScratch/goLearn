package methods

import
(
	"testing"
)


func TestCalcDistanceBetweenTwoPoints(t *testing.T) {
	var expected float64 = 7.0710678118654755
	a := Point{5, 5}
	b := Point{0, 0}
	d := a.Distance(b)

	if d != expected {
		t.Errorf("Distance between {0,0} and {10,10} should be %d, was %d", expected, d)
	}
}