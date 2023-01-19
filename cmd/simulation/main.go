package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/creatures/internal/render"
	"github.com/lukegriffith/creatures/internal/world"
	"github.com/lukegriffith/creatures/internal/worldMap"
)

var (
	generations   = 1
	selectionZone = worldMap.Bounds{
		ID:     0,
		X:      100,
		Y:      100,
		Width:  100,
		Height: 100,
	}
)

func main() {

	// Required for PixelGL to start.
	pixelgl.Run(run)
}

func run() {
	world := world.NewWorld()
	world.Populate(50)
	log.Println(len(world.Qt.GetObjects()))

	for i := 0; i < generations; i++ {
		log.Println("Generation:", i)
		render.Render(world, selectionZone)
		world = world.BreedInSelection(50, selectionZone)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Replay last loop: Y/N ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("Y", text) == 0 {
			win := render.Render(world, selectionZone)
			win.Destroy()
		}

		if strings.Compare("N", text) == 0 {
			os.Exit(0)
		}
	}

}
