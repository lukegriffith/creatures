package creatures

import (
	"github.com/lukegriffith/creatures/internal/neural"
	"github.com/lukegriffith/creatures/internal/util"
	"github.com/lukegriffith/creatures/internal/world"
)

type Attributes struct {
	Health   int
	Strength int
	// Focus is used as vision distance.
	Focus int
	Speed int
	Age   int
}

func NewRandomAttributes() Attributes {
	return Attributes{
		util.RandomInt(60, 100),
		util.RandomInt(1, 10),
		util.RandomInt(10, 25),
		util.RandomInt(1, 4),
		util.RandomInt(0, 20),
	}
}

type Creature struct {
	WorldObjectID world.ObjectID
	WorldMap      *world.Map
	Stats         Attributes
	Brain         neural.Brain
}

func (c Creature) Sense() {

}

func SpawnCreature(wm *world.Map) Creature {
	brain := neural.NewBrain()
	stats := NewRandomAttributes()
	return NewCreature(stats, brain, wm)
}

func BreedCreaturePair(c1 Creature, c2 Creature, wm *world.Map) Creature {
	brain := c1.Brain.Crossover(&c2.Brain)
	attr := Attributes{
		Health:   util.RandomInt(60, 100),
		Strength: (c1.Stats.Strength + c2.Stats.Strength) / 2,
		Focus:    (c1.Stats.Focus + c2.Stats.Focus) / 2,
		Speed:    (c1.Stats.Speed + c2.Stats.Speed) / 2,
		Age:      0,
	}
	return NewCreature(attr, brain, wm)
}

func NewCreature(s Attributes, b neural.Brain, wm *world.Map) Creature {
	return Creature{
		WorldObjectID: wm.AddRandomObject(),
		WorldMap:      wm,
		Stats:         s,
		Brain:         b,
	}
}
