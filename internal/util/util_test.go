package util

import "testing"

func TestRandomNumber(t *testing.T) {
	n1 := RandomInt(10, 100)
	n2 := RandomInt(10, 100)

	t.Log(n1, n2)
	if n1 == n2 {
		t.Fail()
	}

	n1f := RandomFloat(10, 100)
	n2f := RandomFloat(10, 100)

	t.Log(n1f, n2f)
	if n1f == n2f {
		t.Fail()
	}
}

func TestMinMax(t *testing.T) {
	f := []float64{1.0, 1.0, 2.0, 0.5}
	min, max := MinMax(f)
	if min != 3 {
		t.Fail()
	}
	if max != 2 {
		t.Fail()
	}
}
