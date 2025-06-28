package main

import (
	"defend/assets"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	alienVerticalRange = 210.0 - 40.0
	alienAmplitude     = alienVerticalRange / 2
	alienBaseY         = 40.0 + alienAmplitude
	alienFrequency     = (2 * math.Pi) / 320 
)

type Alien struct {
	X         float64
	Y         float64
	Image     *ebiten.Image
	Active    bool
	Facing	FACING 
}

func GetDirectionFromFacing(facing FACING) float64 {
	if facing == RIGHT {
		return 1.0
	}
	return -1.0
}

func NewAlien(x float64) *Alien {
	facing := RIGHT 
	if rand.Float64() < 0.5 {
		facing = LEFT
	}
	// Calculate initial Y based on X
	y := alienBaseY + math.Sin(x*alienFrequency)*alienAmplitude
	return &Alien{
		X:         x,
		Y:         y,
		Image:     assets.AlienSprite,
		Active:    true,
		Facing:    facing,
	}
}

func (a *Alien) Update() {
	speed := 0.75
	a.X += GetDirectionFromFacing(a.Facing) * speed
	a.Y = alienBaseY + math.Sin(a.X*alienFrequency)*alienAmplitude
}

func CheckAlienSpawn(activeAliens []*Alien, terrainWidth float64) []*Alien {
	if len(activeAliens) < 6 {
		randomX := rand.Float64() * terrainWidth
		newAlien := NewAlien(randomX)
		activeAliens = append(activeAliens, newAlien)
	}
	return activeAliens
}

func (a *Alien) Draw(screen *ebiten.Image, camera *Camera, terrainWidth float64) {
	drawX := a.X - camera.X

	// Handle wrapping - if the alien appears to be very far away due to wrapping,
	// adjust its position to appear on the correct side
	if drawX > terrainWidth/2 {
		drawX -= terrainWidth
	} else if drawX < -terrainWidth/2 {
		drawX += terrainWidth
	}

	// Only draw if alien is visible on screen 
	if drawX > -50 && drawX < camera.Width+50 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(drawX, a.Y)
		screen.DrawImage(a.Image, op)
	}
}
