package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type eyes struct {
	pos          pixel.Vec
	color        color.Color
	size         float64
	pupil_offset float64
}

// Draws eyes to screen
func (p *eyes) draw(imd *imdraw.IMDraw) {
	// Balls
	imd.Color = p.color
	imd.Push(p.pos.Add(pixel.V(-(p.size*1.1 + 5), 0)))
	imd.Circle(p.size, 0)
	imd.Push(p.pos.Add(pixel.V((p.size*1.1 + 5), 0)))
	imd.Circle(p.size, 0)
	// Pupils
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(p.pos.Add(pixel.V(-(p.size*1.1 + 5 + p.pupil_offset), 0)))
	imd.Circle(p.size/4, 0)
	imd.Push(p.pos.Add(pixel.V((p.size*1.1 + 5 + p.pupil_offset), 0)))
	imd.Circle(p.size/4, 0)
}

// Posistion to look
func (p *eyes) look(l float64) {
	ball_rad := (p.size*1.1 + 5)
	p.pupil_offset = ((l * (ball_rad - -(ball_rad))) / 100) + -(ball_rad)
}

func (p *eyes) update(imd *imdraw.IMDraw) {

}
