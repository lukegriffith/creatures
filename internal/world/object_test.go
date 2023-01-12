package world

import "testing"

func TestOID(t *testing.T) {
	var oi ObjectIDFactory

	id1 := oi.ID()
	id2 := oi.ID()

	t.Log(id1, id2)

	if !(id1 == 0) {
		t.Fail()
	}

	if !(id2 == 1) {
		t.Fail()
	}
}
