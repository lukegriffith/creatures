package worldMap

import "sync"

var (
	oi ObjectIDFactory
)

type ObjectID int

type Object struct {
	ID     ObjectID
	Bounds *Bounds
}

func NewObject(X float64, Y float64, Width float64, Height float64) Object {
	id := oi.ID()
	return Object{
		id, &Bounds{id, X, Y, Width, Height},
	}
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
