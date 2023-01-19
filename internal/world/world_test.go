package world

import (
	"testing"

	"github.com/lukegriffith/creatures/internal/worldMap"
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

func TestZeroBreeding(t *testing.T) {
	world := NewWorld()
	world.Populate(50)
	world.BreedInSelection(50, worldMap.NewBounds(700, 700, 1, 1))
	objects := world.Qt.GetObjects()
	t.Log(len(objects))
}

func TestNewWorldFromCurrent(t *testing.T) {
	world := NewWorld()
	world.Populate(50)
	world = world.NewWorldFromCreatures()
	for i := 0; i < 100; i++ {
		world.Cycle()
	}
}
