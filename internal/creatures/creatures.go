package creatures

import (
	"log"

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
		Health:   util.RandomFloat(60, 100),
		Strength: util.RandomFloat(1, 10),
		Focus:    util.RandomFloat(10, 25),
		Speed:    util.RandomFloat(1, 3),
		Age:      util.RandomFloat(0, 20),
	}
}

type Creature struct {
	WorldObjectID worldMap.ObjectID
	Qt            *worldMap.Quadtree
	Stats         Attributes
	Brain         neural.Brain
}

func (c Creature) Cycle(object worldMap.Bounds, qt *worldMap.Quadtree, osc float64, age float64) worldMap.Bounds {
	// Sense Environment
	inputs := c.Sense(osc, age)
	// Input to NN
	log.Println("Network In", inputs.ReturnFloatArray())
	outputArr := c.Brain.Network.Predict(inputs.ReturnFloatArray())
	// Process output neurons
	output := neural.MapOutputNeurons(outputArr)
	// Return new bounds
	return c.Move(output, object, qt)
}

func (c Creature) Sense(osc float64, age float64) neural.InputNeurons {
	obj, err := c.Qt.GetObject(c.WorldObjectID)
	if err != nil {
		panic("Cant find creature")
	}
	return neural.MapInputNeurons(obj, c.Qt, c.Stats.Focus, osc, age)

}

func (c Creature) Move(n neural.OutputNeurons, bounds worldMap.Bounds, qt *worldMap.Quadtree) worldMap.Bounds {

	leftIdx, rightIdx, downIdx, upIdx := 0, 1, 2, 3
	_, largestIndex := util.MinMax([]float64{n.MLeft, n.MRight, n.MDown, n.MUp})
	stride := c.Stats.Speed
	bounds, err := qt.GetObject(c.WorldObjectID)
	if err != nil {
		panic("unable to find creature")
	}
	log.Println(n)
	log.Println("IDX:", largestIndex)
	log.Println()
	if largestIndex == leftIdx {
		newX := bounds.X - stride
		bounds.X = newX
	}
	if largestIndex == rightIdx {
		newX := bounds.X + stride
		bounds.X = newX
	}
	if largestIndex == downIdx {
		newY := bounds.Y - stride
		bounds.Y = newY
	}
	if largestIndex == upIdx {
		newY := bounds.Y + stride
		bounds.Y = newY
	}
	return bounds
}

func SpawnCreature(qt *worldMap.Quadtree) Creature {
	brain := neural.NewBrain()
	stats := NewRandomAttributes()
	return NewCreature(stats, brain, qt)
}

func BreedCreaturePair(c1 Creature, c2 Creature, qt *worldMap.Quadtree) Creature {
	brain := c1.Brain.Crossover(&c2.Brain)
	attr := Attributes{
		Health:   util.RandomFloat(60, 100),
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
