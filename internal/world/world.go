package world

import (
	"github.com/lukegriffith/creatures/internal/creatures"
	"github.com/lukegriffith/creatures/internal/worldMap"
)

type World interface {
	Cycle()
	Reset()
	GetObjects() []worldMap.Object
}

type RealTimeWorld struct {
	Map       *worldMap.Map
	nCycle    int
	creatures []creatures.Creature
}

func (w *RealTimeWorld) Populate(pop int) {
	for i := 0; i < pop; i++ {
		w.creatures = append(w.creatures, creatures.SpawnCreature(w.Map))
	}
}

func (w *RealTimeWorld) Cycle() {
	w.nCycle = w.nCycle + 1
	w.Map.ClearMap()
	for _, c := range w.creatures {
		wo, err := w.Map.GetObject(c.WorldObjectID)
		if err != nil {
			panic("creature not found")
		}
		woUpdate := c.Cycle(wo)
		id, err := w.Map.AddObject(woUpdate.X, woUpdate.Y, woUpdate.Width, woUpdate.Head)
		if err != nil {
			/// TODO: Figure out how to handle this. If the object cant be placed in the new
			// location. but also cant fit in the old location where does it go?
			w.Map.AddObject()
		}

		// cycle creature
		// insert into map
	}
}
