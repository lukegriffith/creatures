package world

import (
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
		woUpdate := c.Cycle(wo, w.Qt)
		if worldMap.CheckCollision(woUpdate, w.Qt) {
			// Add to old palce
			err = newQt.AddObject(wo)
			if err != nil {
				log.Panicf("Cant insert object in old location", err)
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
