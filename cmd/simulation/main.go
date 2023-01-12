package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/creatures/internal/creatures"
	"github.com/lukegriffith/creatures/internal/render"
	"github.com/lukegriffith/creatures/internal/world"
)

func main() {
	// Required for PixelGL to start.
	pixelgl.Run(run)
}

func run() {
	world := world.NewMap()
	for i := 1; i < 40; i++ {
		creatures.SpawnCreature(world)
	}
	render.Render(world)
}
