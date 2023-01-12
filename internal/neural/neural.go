package neural

import (
	"github.com/lukegriffith/creatures/internal/util"
	"github.com/patrikeh/go-deep"
)

type Brain struct {
	network *deep.Neural
}

func NewBrain() Brain {
	return Brain{
		CreateNetwork(6, []int{2, 2, 4}),
	}
}

func CreateNetwork(inputs int, layout []int) *deep.Neural {
	return deep.NewNeural(&deep.Config{
		/* Input dimensionality */
		Inputs: inputs,
		/* Two hidden layers consisting of two neurons each, and a single output */
		Layout: layout,
		/* Activation functions: Sigmoid, Tanh, ReLU, Linear */
		Activation: deep.ActivationSigmoid,
		/* Determines output layer activation & loss function:
		ModeRegression: linear outputs with MSE loss
		ModeMultiClass: softmax output with Cross Entropy loss
		ModeMultiLabel: sigmoid output with Cross Entropy loss
		ModeBinary: sigmoid output with binary CE loss */
		Mode: deep.ModeBinary,
		/* Weight initializers: {deep.NewNormal(μ, σ), deep.NewUniform(μ, σ)} */
		Weight: deep.NewNormal(1.0, 0.0),
		/* Apply bias */
		Bias: true,
	})
}

func (n *Brain) Crossover(partner *Brain) Brain {
	crossoverPoint := util.RandomInt(0, len(n.network.Layers))
	l1 := n.network.Layers[:crossoverPoint]
	l2 := partner.network.Layers[crossoverPoint:]
	b1 := n.network.Biases[:crossoverPoint]
	b2 := partner.network.Biases[crossoverPoint:]

	l3 := append(l1, l2...)
	b3 := append(b1, b2...)
	brain := Brain{
		&deep.Neural{
			Layers: l3,
			Biases: b3,
			Config: partner.network.Config,
		},
	}
	return brain
}
