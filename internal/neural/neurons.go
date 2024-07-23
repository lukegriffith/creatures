package neural

import (
	"log"

	"github.com/lukegriffith/creatures/internal/worldMap"
)

type InputNeurons struct {
	PopLeft   float64
	PopRight  float64
	PopDown   float64
	PopUp     float64
	Oscilator float64
	Age       float64
}

func popLeft(b worldMap.Bounds, qt *worldMap.Quadtree, vision float64) float64 {
	bounds := worldMap.Bounds{
		ID:     0,
		X:      b.X - vision,
		Y:      b.Y,
		Width:  6,
		Height: 6,
	}
	insercetions := qt.RetrieveIntersections(bounds)
	return float64(len(insercetions))
}

func popRight(b worldMap.Bounds, qt *worldMap.Quadtree, vision float64) float64 {
	bounds := worldMap.Bounds{
		ID:     0,
		X:      b.X + vision,
		Y:      b.Y,
		Width:  6,
		Height: 6,
	}
	insercetions := qt.RetrieveIntersections(bounds)
	return float64(len(insercetions))
}

func popDown(b worldMap.Bounds, qt *worldMap.Quadtree, vision float64) float64 {
	bounds := worldMap.Bounds{
		ID:     0,
		X:      b.X,
		Y:      b.Y - vision,
		Width:  6,
		Height: 6,
	}
	insercetions := qt.RetrieveIntersections(bounds)
	return float64(len(insercetions))
}

func popUp(b worldMap.Bounds, qt *worldMap.Quadtree, vision float64) float64 {
	bounds := worldMap.Bounds{
		ID:     0,
		X:      b.X,
		Y:      b.Y + vision,
		Width:  6,
		Height: 6,
	}
	insercetions := qt.RetrieveIntersections(bounds)
	return float64(len(insercetions))
}

func MapInputNeurons(b worldMap.Bounds, qt *worldMap.Quadtree, vision float64, osc float64, worldAge float64) InputNeurons {
	return InputNeurons{
		PopLeft:   popLeft(b, qt, vision),
		PopRight:  popRight(b, qt, vision),
		PopDown:   popDown(b, qt, vision),
		PopUp:     popUp(b, qt, vision),
		Oscilator: osc,
		Age:       worldAge,
	}
}

func (n InputNeurons) ReturnFloatArray() []float64 {
	return []float64{
		n.PopLeft,
		n.PopRight,
		n.PopDown,
		n.PopRight,
		n.Oscilator,
		n.Age,
	}
}

type OutputNeurons struct {
	MLeft  float64
	MRight float64
	MDown  float64
	MUp    float64
}

func MapOutputNeurons(in []float64) OutputNeurons {
	inLen := len(in)
	if inLen != 4 {
		panic("input data invalid")
	}
	log.Println("Output:", in)
	return OutputNeurons{
		MLeft:  in[0],
		MRight: in[1],
		MDown:  in[2],
		MUp:    in[3],
	}
}
