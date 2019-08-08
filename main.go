package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// -- To Do --
// clean up pixel run setup.

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "CV Monster",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
		//Undecorated: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	eye := &eyes{
		pos:   pixel.V(512, 384),
		color: pixel.RGB(.45, .49, .60),
		size:  100,
	}

	imd := imdraw.New(nil)

	track := newTracker()

	eye.look(80)
	fps := time.Tick(time.Second / 10)
	for !win.Closed() {
		imd.Clear()
		eye.draw(imd)
		win.Clear(pixel.RGB(1, 1, 1))
		imd.Draw(win)
		win.Update()
		track.processFrame()
		<-fps
	}
}

func main() {
	//pixelgl.Run(run)
	track := newTracker()
	for {
		track.processFrame()
	}

}
