package world

import (
	"testing"
)

func TestWorldPopulation(t *testing.T) {
	world := NewWorld()
	world.Populate(50)
	pop := len(world.Qt.GetObjects())
	if pop < 50 {
		t.Log(pop)
	}

	for i := 0; i < 100; i++ {
		world.Cycle()
	}
}
