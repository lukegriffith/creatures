package worldMap

import "sync"

var (
	oi ObjectIDFactory
)

type ObjectID int

func NewBounds(X float64, Y float64, Width float64, Height float64) Bounds {
	id := oi.ID()
	return Bounds{id, X, Y, Width, Height}
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
