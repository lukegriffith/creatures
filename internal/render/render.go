package render

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lukegriffith/creatures/internal/worldMap"
	"golang.org/x/image/colornames"
)

const (
	frameRate time.Duration = 33 * time.Millisecond
	simLength               = 500
)

func renderCreature(x float64, y float64, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Clear()
	imd.Color = colornames.Navy
	imd.Push(pixel.V(x, y))
	imd.Ellipse(pixel.V(5, 5), 0)
	imd.Draw(win)
}

func Render(wm *worldMap.Map) {
	var cycle int
	cycle = 0

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
	for !win.Closed() {

		select {
		case <-tick:
			win.Clear(colornames.Aliceblue)

			for _, obj := range wm.GetObjects() {
				fmt.Println(obj)
				renderCreature(obj.X, obj.Y, imd, win)
			}
		}

		if cycle < 498 {
			cycle++
		}

		win.Update()
	}
}
