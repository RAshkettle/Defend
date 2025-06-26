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

// Movement constants
const (
	PLAYER_MOVE_SPEED = 0.1 
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
	screenWidth := int(camera.Width)
	minimapWidth := screenWidth / 3
	minimapX := (screenWidth - minimapWidth) / 2
	leftEdge := float64(minimapX)
	rightEdge := float64(minimapX + minimapWidth)
	minY := 35.0 
	maxY := float64(camera.Height - 10)

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		camera.MoveLeft(terrainWidth)
		p.Facing = LEFT
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		camera.MoveRight(terrainWidth)
		p.Facing = RIGHT
	}



	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 2
		if p.Y < minY {
			p.Y = minY
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 2
		if p.Y > maxY {
			p.Y = maxY
		}
	}

	targetX := leftEdge
	if p.Facing == LEFT {
		targetX = rightEdge
	}

	distToTarget := targetX - p.X
	p.X += distToTarget * PLAYER_MOVE_SPEED
	
	return nil
}
