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

func (p *Player) Update(camera *Camera, terrainWidth float64) error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA){
		camera.MoveLeft(terrainWidth)
		p.Facing = LEFT
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD){
		camera.MoveRight(terrainWidth)
		p.Facing = RIGHT
	}

	return nil
}
