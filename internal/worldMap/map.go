package worldMap

import (
	"errors"

	"github.com/lukegriffith/creatures/internal/util"
)

var (
	width  = 500
	height = 500
	size   = float64(2)
)

func NewQuadTree() *Quadtree {
	return &Quadtree{
		Bounds: Bounds{
			X:      0,
			Y:      0,
			Width:  float64(width),
			Height: float64(height),
		},
		MaxObjects: 1,
		MaxLevels:  4,
		Level:      0,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0),
	}
}

func (qt *Quadtree) AddObject(bounds Bounds) error {
	// Does object exist within the map range.
	if !bounds.Intersects(qt.Bounds) {
		return errors.New("object is not within map")
	}
	// does object collide with another.
	if CheckCollision(bounds, qt) {
		return errors.New("collides with another")
	}
	qt.Insert(bounds)
	return nil
}

func CheckCollision(bound Bounds, qt *Quadtree) bool {
	objects := qt.RetrieveIntersections(bound)
	return len(objects) > 0
}

func (qt *Quadtree) AddRandomObject() ObjectID {
	placed := false
	var x, y float64
	var obj Bounds
	var err error

	for !placed {
		x = util.RandomFloat(0, width)
		y = util.RandomFloat(0, height)
		obj = NewBounds(x, y, size, size)
		err = qt.AddObject(obj)
		if err == nil {
			placed = true
		}
	}
	return obj.ID
}

func (qt *Quadtree) GetObject(ID ObjectID) (Bounds, error) {
	for _, obj := range qt.GetObjects() {
		if obj.ID == ID {
			return obj, nil
		}
	}
	return Bounds{0, 0, 0, 0, 0}, errors.New("unable to located bounds by ID")
}

func (qt *Quadtree) GetObjects() []Bounds {
	return qt.Retrieve(qt.Bounds)
}
