package main

import (
	"log"

	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/creatures/internal/render"
	"github.com/lukegriffith/creatures/internal/world"
)

func main() {
	// Required for PixelGL to start.
	pixelgl.Run(run)
}

func run() {
	world := world.NewWorld()
	world.Populate(50)
	log.Println(len(world.Qt.GetObjects()))
	render.Render(world)
}
