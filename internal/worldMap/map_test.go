package worldMap

import (
	"testing"
)

func TestAddingRandomsToMap(t *testing.T) {

	worldMap := NewMap()
	for i := 1; i < 5; i++ {
		obj := worldMap.AddRandomObject()
		t.Log(obj)
	}
}

func TestCollisionByPoint(t *testing.T) {
	worldMap := NewMap()
	_, _ = worldMap.AddObject(10, 10, 2, 2)
	_, err := worldMap.AddObject(10, 10, 2, 2)
	if err != nil {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeCatches(t *testing.T) {
	worldMap := NewMap()
	_, _ = worldMap.AddObject(10, 10, 2, 2)
	_, err := worldMap.AddObject(10, 12, 2, 2)
	if err != nil {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeInserts(t *testing.T) {
	worldMap := NewMap()
	_, _ = worldMap.AddObject(10, 10, 2, 2)
	_, err := worldMap.AddObject(10, 13, 2, 2)
	if err == nil {
		t.Log("no collision as expected")
		return
	}
	t.Fail()
}

func TestOutOfBounds(t *testing.T) {
	worldMap := NewMap()
	_, err := worldMap.AddObject(7000, 7000, 2, 2)
	if err == nil {
		t.Log("object placed when out of bounds")
		t.Fail()
	}
	objects := worldMap.GetObjects()
	if len(objects) > 0 {
		t.Log("object in object list when it shouldn't exist")
		t.Fail()
	}

}
