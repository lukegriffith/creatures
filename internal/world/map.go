package world

import (
	"errors"

	quadtree "github.com/JamesLMilner/quadtree-go"
	"github.com/lukegriffith/creatures/internal/util"
)

var (
	width  = 500
	height = 500
	size   = float64(2)
)

type Map struct {
	width   float64
	height  float64
	objects []Object
	qt      *quadtree.Quadtree
}

func NewMap() *Map {
	qt := &quadtree.Quadtree{
		Bounds: quadtree.Bounds{
			X:      0,
			Y:      0,
			Width:  float64(width),
			Height: float64(height),
		},
		MaxObjects: 1,
		MaxLevels:  4,
		Level:      0,
		Objects:    make([]quadtree.Bounds, 0),
		Nodes:      make([]quadtree.Quadtree, 0),
	}

	return &Map{
		width:   float64(width),
		height:  float64(height),
		objects: make([]Object, 0),
		qt:      qt,
	}
}

func (m *Map) AddObject(X float64, Y float64, Width float64, Height float64) ObjectID {
	bounds := quadtree.Bounds{
		X:      X,
		Y:      Y,
		Width:  Width,
		Height: Height,
	}
	obj, id := NewObject(X, Y, size, size)
	m.qt.Insert(bounds)
	m.objects = append(m.objects, obj)
	return id
}

func checkCollision(bound quadtree.Bounds, m *Map) bool {
	objects := m.qt.RetrieveIntersections(bound)
	return len(objects) > 0
}

func (m *Map) AddRandomObject() ObjectID {
	placed := false
	var bounds quadtree.Bounds
	var x, y float64

	for !placed {
		x = util.RandomFloat(0, width)
		y = util.RandomFloat(0, height)
		bounds = quadtree.Bounds{
			X:      x,
			Y:      y,
			Width:  size,
			Height: size,
		}
		if checkCollision(bounds, m) {
			continue
		}
		placed = true
	}
	return m.AddObject(x, y, size, size)
}

func (m *Map) GetObject(ID ObjectID) (Object, error) {
	for _, obj := range m.objects {
		if obj.ID == ID {
			return obj, nil
		}
	}
	return Object{0, 0, 0, 0, 0}, errors.New("unable to located object by ID")
}

func (m *Map) GetObjects() []Object {
	return m.objects
}
