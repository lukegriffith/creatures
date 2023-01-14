package world

import (
	"github.com/lukegriffith/creatures/internal/creatures"
	"github.com/lukegriffith/creatures/internal/worldMap"
)

type RealTimeWorld struct {
	Qt        *worldMap.Quadtree
	nCycle    int
	creatures []creatures.Creature
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
		woUpdate := c.Cycle(wo)
		if worldMap.CheckCollision(woUpdate, w.Qt) {
			// Add to old palce
			err = newQt.AddObject(wo)
			if err != nil {
				panic("Cant insert object in old location")
			}
		} else {
			// Addd to new place
			err = newQt.AddObject(woUpdate)
			panic("Cant insert object in new location")
		}
	}
	w.Qt = newQt
}
