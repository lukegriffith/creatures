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

func TestBreeding(t *testing.T) {
	world := NewWorld()
	world.Populate(50)
	world.BreedInSelection(50, world.Qt.Bounds)
}

func TestNewWorldFromCurrent(t *testing.T) {
	world := NewWorld()
	world.Populate(50)
	world = world.NewWorldFromCreatures()
	for i := 0; i < 100; i++ {
		world.Cycle()
	}
}
