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

type Map struct {
	objects []Object
	qt      *Quadtree
}

func NewMap() *Map {
	qt := &Quadtree{
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

	return &Map{
		width:   float64(width),
		height:  float64(height),
		objects: make([]Object, 0),
		qt:      qt,
	}
}

func (m *Map) AddObject(obj Object) error {
	bounds := *obj.Bounds
	// Does object exist within the map.
	if !bounds.Intersects(m.qt.Bounds) {
		return errors.New("object is not within map")
	}
	if checkCollision(bounds, m) {
		return errors.New("collides with another")
	}

	m.qt.Insert(bounds)
	m.objects = append(m.objects, obj)
	return nil
}

func checkCollision(bound Bounds, m *Map) bool {
	objects := m.qt.RetrieveIntersections(bound)
	return len(objects) > 0
}

func (m *Map) AddRandomObject() ObjectID {
	placed := false
	var x, y float64
	var obj Object
	var err error

	for !placed {
		x = util.RandomFloat(0, width)
		y = util.RandomFloat(0, height)
		obj = NewObject(x, y, size, size)
		err = m.AddObject(obj)
		if err == nil {
			placed = true
		}
	}
	return obj.ID
}

func (m *Map) GetObject(ID ObjectID) (Object, error) {
	for _, obj := range m.objects {
		if obj.ID == ID {
			return obj, nil
		}
	}
	return Object{0, nil}, errors.New("unable to located object by ID")
}

func (m *Map) GetObjects() []Object {
	return m.objects
}

func (m *Map) ClearMap() {
	m.qt.Clear()
}
