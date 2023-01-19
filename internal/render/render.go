package render

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/creatures/internal/world"
	"github.com/lukegriffith/creatures/internal/worldMap"
	"golang.org/x/image/colornames"
)

const (
	frameRate time.Duration = 33 * time.Millisecond
	simLength               = 500
)

/*
imd.Color = colornames.Navy
imd.Push(pixel.V(200, 500), pixel.V(800, 500))
imd.Ellipse(pixel.V(120, 80), 0)
*/
func renderCreature(x float64, y float64, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Clear()
	imd.Color = colornames.Navy
	imd.Push(pixel.V(x, y))
	imd.Ellipse(pixel.V(3, 3), 0)
	imd.Draw(win)
}

func renderCreatures(c []worldMap.Bounds, imd *imdraw.IMDraw) {

	positions := make([]pixel.Vec, 0)
	for _, b := range c {
		positions = append(positions, pixel.V(b.X, b.Y))
	}
	imd.Color = colornames.Navy
	imd.Push(positions...)
	imd.Ellipse(pixel.V(3, 3), 0)
}

func renderSelectionZone(x float64, y float64, width float64, height float64, imd *imdraw.IMDraw) {
	rect := pixel.R(x, y, x+width, y+height)
	imd.Color = colornames.Red
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(1)

}
func Render(w *world.RealTimeWorld, selection worldMap.Bounds) *pixelgl.Window {
	cycle := 0

	cfg := pixelgl.WindowConfig{
		Title:  "Creatures",
		Bounds: pixel.R(0, 0, 500, 500),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Aliceblue)

	tick := time.Tick(frameRate)

	imd := imdraw.New(nil)
	log.Println("creatures in world: ", len(w.Qt.GetObjects()))
	for cycle < 100 && !win.Closed() {

		select {
		case <-tick:
			win.Clear(colornames.Aliceblue)

			renderCreatures(w.Qt.GetObjects(), imd)
			renderSelectionZone(selection.X, selection.Y, selection.Width, selection.Height, imd)

			imd.Draw(win)
			imd.Clear()
		}

		if cycle < 100 {
			w.Cycle()
			cycle++
		}

		win.Update()
	}
	return win
}
