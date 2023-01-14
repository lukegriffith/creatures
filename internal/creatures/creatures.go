package creatures

import (
	"github.com/lukegriffith/creatures/internal/neural"
	"github.com/lukegriffith/creatures/internal/util"
	"github.com/lukegriffith/creatures/internal/worldMap"
)

type Attributes struct {
	Health   float64
	Strength float64
	// Focus is used as vision distance.
	Focus float64
	Speed float64
	Age   float64
}

func NewRandomAttributes() Attributes {
	return Attributes{
		util.RandomFloat(60, 100),
		util.RandomFloat(1, 10),
		util.RandomFloat(10, 25),
		util.RandomFloat(1, 4),
		util.RandomFloat(0, 20),
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
	inputs := c.Sense()
	// Input to NN
	outputArr := c.Brain.Network.Predict(inputs.ReturnFloatArray())
	// Process output neurons
	output := neural.MapOutputNeurons(output)
	// Return new bounds

}

func (c Creature) Sense() neural.InputNeurons {
	// TO BE IMPLEMENTED
	obj, err := c.Qt.GetObject(c.WorldObjectID)
	if err != nil {
		panic("Cant find creature")
	}
	return neural.MapInputNeurons(obj, c.Qt, c)

}

func (c Creature) Move(n neural.OutputNeurons) worldMap.Bounds {
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
