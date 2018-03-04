package methods

import
(
	"testing"
)


func TestCalcDistanceBetweenTwoPoints(t *testing.T) {
	var expected float64 = 5
	a := Point{1, 1}
	b := Point{5, 4}
	d := a.Distance(b)

	if d != expected {
		t.Errorf("Distance between {0,0} and {10,10} should be %d, was %d", expected, d)
	}
}