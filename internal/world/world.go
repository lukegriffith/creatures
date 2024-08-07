package world

import (
	"errors"
	"log"

	"github.com/lukegriffith/creatures/internal/creatures"
	"github.com/lukegriffith/creatures/internal/worldMap"
)

type RealTimeWorld struct {
	Qt        *worldMap.Quadtree
	nCycle    int
	creatures []creatures.Creature
}

func NewWorld() *RealTimeWorld {
	return &RealTimeWorld{
		Qt:        worldMap.NewQuadTree(),
		nCycle:    0,
		creatures: nil,
	}
}

func (w *RealTimeWorld) Populate(pop int) {
	for i := 0; i < pop; i++ {
		w.creatures = append(w.creatures, creatures.SpawnCreature(w.Qt))
	}
}

func (w *RealTimeWorld) Cycle() {
	w.nCycle = w.nCycle + 1
	newQt := worldMap.NewQuadTree()
	for _, c := range w.creatures {
		wo, err := w.Qt.GetObject(c.WorldObjectID)
		if err != nil {
			panic("creature not found")
		}
		woUpdate := c.Cycle(wo, w.Qt, w.Oscilator(), float64(w.GetCycle()))
		if worldMap.CheckCollision(woUpdate, w.Qt) {
			// Add to old palce
			err = newQt.AddObject(wo)
			if err != nil {
				log.Panic("Cant insert object in old location", err)
			}
		} else {
			// Addd to new place
			err = newQt.AddObject(woUpdate)
			if err != nil {
				err = newQt.AddObject(wo)
				if err != nil {
					log.Println("Cant insert object in new or old location", err)
				}
			}
		}
	}
	w.Qt = newQt
}

func (w *RealTimeWorld) GetCreaturesInList(b []worldMap.Bounds) []creatures.Creature {
	c := make([]creatures.Creature, 0)
	cmap := map[worldMap.ObjectID]creatures.Creature{}

	for _, creature := range w.creatures {
		cmap[creature.WorldObjectID] = creature
	}
	for _, bounds := range b {
		c = append(c, cmap[bounds.ID])
	}
	return c
}

func (w *RealTimeWorld) BreedInSelection(n int, b worldMap.Bounds) (*RealTimeWorld, error) {
	selectedBounds := w.Qt.RetrieveIntersections(b)
	selectedCreatures := w.GetCreaturesInList(selectedBounds)
	newPop := make([]creatures.Creature, 0)
	newQt := worldMap.NewQuadTree()
	popCount := 0

	if len(selectedBounds) < 2 {
		return nil, errors.New("population too sparse")
	}

out:
	for {
		for _, c := range selectedCreatures {
			for _, c2 := range selectedCreatures {
				newPop = append(newPop, creatures.BreedCreaturePair(c, c2, newQt))
				popCount++
				if popCount >= n {
					break out
				}
			}
		}
	}

	return &RealTimeWorld{
		Qt:        newQt,
		nCycle:    0,
		creatures: newPop,
	}, nil
}

func (w *RealTimeWorld) NewWorldFromCreatures() *RealTimeWorld {
	newW := &RealTimeWorld{
		Qt:        worldMap.NewQuadTree(),
		nCycle:    0,
		creatures: nil,
	}
	for _, c := range w.creatures {
		id := newW.Qt.AddRandomObject()
		c.WorldObjectID = id
		c.Qt = newW.Qt
		newW.creatures = append(newW.creatures, c)
	}
	return newW
}

func (w *RealTimeWorld) Oscilator() float64 {
	return float64(w.nCycle%10) / 10
}

func (w *RealTimeWorld) GetCycle() int {
	return w.nCycle
}
