package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Terrain struct {
	points       []float64
	width        float64
	terrainColor color.RGBA
}

func NewTerrain(screenWidth float64) *Terrain {
	totalWidth := screenWidth * 3
	numPoints := int(totalWidth / 5)

	points := make([]float64, numPoints)

	for i := 0; i < numPoints; i++ {
		x := float64(i) * 5.0

		baseHeight := 40.0

		height := baseHeight +
			20.0*math.Sin(x*0.01) +
			15.0*math.Sin(x*0.02+0.5) +
			25.0*math.Sin(x*0.005-0.7) +
			10.0*math.Max(0, math.Sin(x*0.03+1.2))

		points[i] = height
	}

	return &Terrain{
		points:       points,
		width:        totalWidth,
		terrainColor: color.RGBA{255, 140, 0, 255},
	}
}

func (t *Terrain) Draw(screen *ebiten.Image, camera *Camera) {
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())

	cameraX := camera.X
	for cameraX < 0 {
		cameraX += t.width
	}
	for cameraX >= t.width {
		cameraX -= t.width
	}

	pointSpacing := 5.

	for screenX := 0.; screenX < screenWidth; screenX += pointSpacing {
		worldX := screenX + cameraX

		for worldX >= t.width {
			worldX -= t.width
		}

		index := int(worldX / pointSpacing)
		nextIndex := (index + 1) % len(t.points)

		if index >= 0 && index < len(t.points) {
			height := t.points[index]
			nextHeight := t.points[nextIndex]

			y1 := screenHeight - height
			y2 := screenHeight - nextHeight
			x1 := screenX
			x2 := screenX + pointSpacing

			vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, t.terrainColor, false)
		}
	}
}
