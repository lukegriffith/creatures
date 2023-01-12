package neural

import (
	"testing"
)

func TestCrossoverCreature(t *testing.T) {
	b1 := &Brain{
		CreateNetwork(6, []int{2, 2, 4}),
	}
	b2 := &Brain{
		CreateNetwork(6, []int{2, 2, 4}),
	}

	b3 := b1.Crossover(b2)
	t.Log(b3)

	output := b3.network.Predict([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0})
	t.Log(output)
}
