package worldMap

import "sync"

var (
	oi ObjectIDFactory
)

type ObjectID int

type Object struct {
	ID     ObjectID
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewObject(X float64, Y float64, Width float64, Height float64) (Object, ObjectID) {
	id := oi.ID()
	return Object{
		id, X, Y, Width, Height,
	}, id

}

type ObjectIDFactory struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         ObjectID
}

func (o *ObjectIDFactory) ID() (id ObjectID) {
	o.Lock()
	defer o.Unlock()
	id = o.id
	o.id++
	return id
}
