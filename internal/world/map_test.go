package world

import (
	"testing"

	"github.com/JamesLMilner/quadtree-go"
)

func TestAddingRandomsToMap(t *testing.T) {

	worldMap := NewMap()
	for i := 1; i < 5; i++ {
		id := worldMap.AddRandomObject()
		t.Log(id)
		obj, _ := worldMap.GetObject(id)
		t.Log(obj)
	}
}

func TestCollisionByPoint(t *testing.T) {
	worldMap := NewMap()
	worldMap.AddObject(10, 10, 2, 2)
	b := quadtree.Bounds{10, 10, 2, 2}
	result := checkCollision(b, worldMap)

	if result {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeCatches(t *testing.T) {
	worldMap := NewMap()
	worldMap.AddObject(10, 10, 2, 2)
	b := quadtree.Bounds{10, 12, 2, 2}
	result := checkCollision(b, worldMap)

	if result {
		t.Log("collision as expected")
		return
	}
	t.Fail()
}

func TestCollisionBySizeInserts(t *testing.T) {
	worldMap := NewMap()
	worldMap.AddObject(10, 10, 2, 2)
	b := quadtree.Bounds{10, 13, 2, 2}
	result := checkCollision(b, worldMap)

	if !result {
		t.Log("no collision as expected")
		return
	}
	t.Fail()
}
