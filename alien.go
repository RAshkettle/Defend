package main

import (
	"defend/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Alien struct {
	X      float64
	Y      float64
	Image  *ebiten.Image
	Active bool
}

func NewAlien(x, y float64) *Alien {
	return &Alien{
		X:      x,
		Y:      y,
		Image:  assets.AlienSprite,
		Active: true,
	}
}

func CheckAlienSpawn(activeAliens []*Alien, terrainWidth float64) []*Alien {
	if len(activeAliens) < 6 {
		randomX := rand.Float64() * terrainWidth
		minY := 40.0
		maxY := 210.0
		randomY := minY + rand.Float64()*(maxY-minY)
		newAlien := NewAlien(randomX, randomY)
		activeAliens = append(activeAliens, newAlien)
	}
	return activeAliens
}

func (a *Alien) Draw(screen *ebiten.Image, camera *Camera, terrainWidth float64) {
	// Calculate alien position relative to camera, handling wrapping
	drawX := a.X - camera.X

	// Handle wrapping - if the alien appears to be very far away due to wrapping,
	// adjust its position to appear on the correct side
	if drawX > terrainWidth/2 {
		drawX -= terrainWidth
	} else if drawX < -terrainWidth/2 {
		drawX += terrainWidth
	}

	// Only draw if alien is visible on screen (with some margin)
	if drawX > -50 && drawX < camera.Width+50 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(drawX, a.Y)
		screen.DrawImage(a.Image, op)
	}
}
