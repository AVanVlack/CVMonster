package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type eyes struct {
	pos   pixel.Vec
	color color.Color
	size  float64
}

func (p *eyes) draw(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(p.pos.Add(pixel.V(-(p.size*1.1 + 5), 0)))
	imd.Circle(p.size, 0)
	imd.Push(p.pos.Add(pixel.V((p.size*1.1 + 5), 0)))
	imd.Circle(p.size, 0)
}

func (p *eyes) look(imd *imdraw.IMDraw) {

}

func (p *eyes) update(imd *imdraw.IMDraw) {

}
