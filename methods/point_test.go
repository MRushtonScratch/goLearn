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

func TestCanScaleAPoint(t *testing.T) {
	expected := &Point{2,4}
	p := Point{1,2}

	p.ScaleBy(2)

	if p.X != expected.X || p.Y != expected.Y {
		t.Errorf("Point did not scale %s - %s", p, expected)
	}
}