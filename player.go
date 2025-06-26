package main

import (
	"defend/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type FACING int

const (
	RIGHT FACING = iota
	LEFT
)

type Player struct {
	X      float64
	Y      float64
	Image  *ebiten.Image
	Facing FACING
}

func NewPlayer() *Player {
	return &Player{
		Image:  assets.PlayerSprite,
		X:      150.0,
		Y:      100.0,
		Facing: RIGHT,
	}
}
