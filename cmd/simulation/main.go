package main

import (
	"bufio"
	"context"
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
	generations = 25
	_           = worldMap.Bounds{
		ID:     0,
		X:      0,
		Y:      0,
		Width:  100,
		Height: 500,
	}
	selectionZone = worldMap.Bounds{
		ID:     0,
		X:      400,
		Y:      0,
		Width:  100,
		Height: 500,
	}
	ctx = context.Background()
)

func main() {
	pixelgl.Run(run)
}

/*
func run() {
	world := ctx.Value("world").(*world.RealTimeWorld)
	selectionZone := ctx.Value("selectionZone").(worldMap.Bounds)
	win := render.Render(world, selectionZone)
	win.Destroy()
}

*/

func run() {
	world := world.NewWorld()
	world.Populate(100)
	log.Println(len(world.Qt.GetObjects()))
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < generations; i++ {

		for _, b := range world.Qt.GetObjects() {
			log.Printf("ID %d X %.0f Y %.0f", b.ID, b.X, b.Y)
		}
		//log.Println("press enter")
		//_, _ = reader.ReadString('\n')
		log.Println("Generation:", i)
		/*
			ctx = context.WithValue(ctx, "world", world)
			ctx = context.WithValue(ctx, "selectionZone", selectionZone)
		*/
		//pixelgl.Run(run)

		win := render.Render(world, selectionZone)
		win.Destroy()

		log.Println("Breeding Fittest")
		world = world.BreedInSelection(50, selectionZone)
		log.Println("Done")
	}

	for {
		fmt.Print("Replay last loop: Y/N ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("Y", text) == 0 {
			world = world.NewWorldFromCreatures()

			/*
				ctx = context.WithValue(ctx, "world", world)
				ctx = context.WithValue(ctx, "selectionZone", selectionZone)
			*/
			//pixelgl.Run(run)
			win := render.Render(world, selectionZone)
			win.Destroy()

		}

		if strings.Compare("N", text) == 0 {
			os.Exit(0)
		}
	}
	// Required for PixelGL to start.

}
