package worldMap

import (
	"testing"
)

func TestAddingRandomsToMap(t *testing.T) {

	worldMap := NewQuadTree()
	for i := 1; i < 5; i++ {
		obj := worldMap.AddRandomObject()
		t.Log(obj)
	}
}

func TestCollisionByPoint(t *testing.T) {
	worldMap := NewQuadTree()
	bounds := NewBounds(10, 10, 2, 2)
	_ = worldMap.AddObject(bounds)
	bounds2 := NewBounds(10, 10, 2, 2)
	err := worldMap.AddObject(bounds2)
	if err != nil {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeCatches(t *testing.T) {
	worldMap := NewQuadTree()
	bounds := NewBounds(10, 10, 2, 2)
	_ = worldMap.AddObject(bounds)
	bounds2 := NewBounds(10, 12, 2, 2)
	err := worldMap.AddObject(bounds2)
	if err != nil {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeInserts(t *testing.T) {
	worldMap := NewQuadTree()
	bounds := NewBounds(10, 10, 2, 2)
	_ = worldMap.AddObject(bounds)
	bounds2 := NewBounds(10, 13, 2, 2)
	err := worldMap.AddObject(bounds2)
	if err == nil {
		t.Log("no collision as expected")
		return
	}
	t.Fail()
}

func TestOutOfBounds(t *testing.T) {
	worldMap := NewQuadTree()
	bounds := NewBounds(7000, 7000, 2, 2)
	err := worldMap.AddObject(bounds)
	if err == nil {
		t.Log("object placed when out of bounds")
		t.Fail()
	}
	if len(worldMap.Objects) > 0 {
		t.Log("object in object list when it shouldn't exist")
		t.Fail()
	}

}
