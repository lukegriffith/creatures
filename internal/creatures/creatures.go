package creatures

import (
	"github.com/lukegriffith/creatures/internal/neural"
	"github.com/lukegriffith/creatures/internal/util"
	"github.com/lukegriffith/creatures/internal/worldMap"
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
	WorldObjectID worldMap.ObjectID
	Qt            *worldMap.Quadtree
	Stats         Attributes
	Brain         neural.Brain
}

func (c Creature) Cycle(object worldMap.Bounds) worldMap.Bounds {
	// Sense Environment
	// Input to NN
	// Process output neurons
	// Return new bounds
}

func (c Creature) Sense() {
	// TO BE IMPLEMENTED
}

func SpawnCreature(qt *worldMap.Quadtree) Creature {
	brain := neural.NewBrain()
	stats := NewRandomAttributes()
	return NewCreature(stats, brain, qt)
}

func BreedCreaturePair(c1 Creature, c2 Creature, qt *worldMap.Quadtree) Creature {
	brain := c1.Brain.Crossover(&c2.Brain)
	attr := Attributes{
		Health:   util.RandomInt(60, 100),
		Strength: (c1.Stats.Strength + c2.Stats.Strength) / 2,
		Focus:    (c1.Stats.Focus + c2.Stats.Focus) / 2,
		Speed:    (c1.Stats.Speed + c2.Stats.Speed) / 2,
		Age:      0,
	}
	return NewCreature(attr, brain, qt)
}

func NewCreature(s Attributes, b neural.Brain, qt *worldMap.Quadtree) Creature {
	return Creature{
		WorldObjectID: qt.AddRandomObject(),
		Qt:            qt,
		Stats:         s,
		Brain:         b,
	}
}
