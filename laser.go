package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Laser struct {
	X                 float64
	Y                 float64
	CurrentLength     int
	Direction         FACING
	DistanceTravelled float64
	Active            bool
}

func NewLaser(x, y float64, facing FACING) *Laser {
	return &Laser{
		X:             x,
		Y:             y,
		CurrentLength: 0,
		Direction:     facing,
		Active:        true,
	}
}

func (l *Laser) Update() {
	if l.Direction == LEFT {
		l.X -= 6
	} else {
		l.X += 6
	}
	if l.CurrentLength < 64 {
		l.CurrentLength += 3
	}
	l.DistanceTravelled += 6
	if l.DistanceTravelled > 150 {
		l.Active = false
	}
}

func (l *Laser) Draw(screen *ebiten.Image, camera *Camera) {
	drawX := l.X - camera.X
	whiteColor := color.RGBA{255, 255, 255, 255}
	lightBlueColor := color.RGBA{200, 220, 255, 255}

	for i := 0; i < l.CurrentLength; i++ {
		if rand.Float64() < 0.2 && i > 2 {
			continue
		}
		pixelX := drawX
		if l.Direction == RIGHT {
			pixelX += float64(i)
		} else {
			pixelX -= float64(i)
		}
		pixelColor := whiteColor
		if rand.Float64() < 0.2 {
			pixelColor = lightBlueColor
		}
		vector.StrokeLine(
			screen,
			float32(pixelX),
			float32(l.Y),
			float32(pixelX+1),
			float32(l.Y),
			1,
			pixelColor,
			false,
		)
	}
}
