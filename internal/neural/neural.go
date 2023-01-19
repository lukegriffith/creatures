package neural

import (
	"github.com/lukegriffith/creatures/internal/util"
	"github.com/patrikeh/go-deep"
)

type Brain struct {
	Network *deep.Neural
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

// TODO Add mutation!
func (n *Brain) Crossover(partner *Brain) Brain {
	crossoverPoint := util.RandomInt(0, len(n.Network.Layers))
	l1 := n.Network.Layers[:crossoverPoint]
	l2 := partner.Network.Layers[crossoverPoint:]
	b1 := n.Network.Biases[:crossoverPoint]
	b2 := partner.Network.Biases[crossoverPoint:]

	l3 := append(l1, l2...)
	b3 := append(b1, b2...)
	brain := Brain{
		&deep.Neural{
			Layers: l3,
			Biases: b3,
			Config: partner.Network.Config,
		},
	}
	return brain
}
