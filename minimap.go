package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Minimap struct {
	terrain       *Terrain
	camera        *Camera
	width         int
	height        int
	posX          int
	posY          int
	terrainColor  color.RGBA
	viewportColor color.RGBA
	borderColor   color.RGBA
}

func NewMinimap(terrain *Terrain, camera *Camera, screenWidth, screenHeight int) *Minimap {
	minimapWidth := screenWidth / 3
	minimapHeight := 20
	minimapX := (screenWidth - minimapWidth) / 2
	minimapY := 5

	return &Minimap{
		terrain:       terrain,
		camera:        camera,
		width:         minimapWidth,
		height:        minimapHeight,
		posX:          minimapX,
		posY:          minimapY,
		terrainColor:  color.RGBA{255, 140, 0, 255},   // Orange terrain
		viewportColor: color.RGBA{255, 255, 255, 255}, // White viewport indicator
		borderColor:   color.RGBA{0, 0, 255, 255},     // Blue border
	}
}

func (m *Minimap) Draw(screen *ebiten.Image) {
	// Draw blue border around minimap
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX+m.width), float32(m.posY), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY+m.height), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX+m.width), float32(m.posY), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)

	// Calculate dimensions and scaling
	terrainWidth := m.terrain.width
	screenWidth := float64(screen.Bounds().Dx())

	// Calculate scaling factors for minimap
	scaleX := float64(m.width) / terrainWidth

	// Find max height for Y-scaling
	maxHeight := 0.0
	for _, h := range m.terrain.points {
		if h > maxHeight {
			maxHeight = h
		}
	}
	scaleY := float64(m.height) / (maxHeight * 1.2) // Scale to fit with some margin

	// Get normalized camera position
	cameraX := m.camera.X
	for cameraX < 0 {
		cameraX += terrainWidth
	}
	for cameraX >= terrainWidth {
		cameraX -= terrainWidth
	}

	// Calculate viewport width in minimap space - represents one screen's worth
	viewportWidth := int((screenWidth / terrainWidth) * float64(m.width))

	// Position viewport indicator based on camera position
	viewportPosition := float64(cameraX) / terrainWidth
	viewportX := m.posX + int(float64(m.width)*viewportPosition) - viewportWidth/2

	// Ensure viewport indicator wraps properly
	if viewportX < m.posX {
		viewportX += m.width
	}
	if viewportX+viewportWidth > m.posX+m.width {
		viewportX -= m.width
	}

	// Draw all terrain points using same pointSpacing as in terrain.Draw
	pointSpacing := 5.0
	lastX := -1.0
	lastY := -1.0

	// Draw all terrain points for all 3 screens
	for i := 0; i < len(m.terrain.points); i++ {
		// Convert world position to minimap position
		worldX := float64(i) * pointSpacing
		minimapX := m.posX + int(worldX*scaleX)%m.width

		// Calculate height at this point
		height := m.terrain.points[i]
		minimapY := m.posY + m.height - int(height*scaleY)

		// Ensure Y stays within minimap bounds
		if minimapY < m.posY {
			minimapY = m.posY
		}

		// Draw line segment if not the first point and not wrapping around edge
		if lastX >= 0 {
			// Check if points are close enough to be connected (avoid wrap lines)
			if abs(minimapX-int(lastX)) < m.width/2 {
				vector.StrokeLine(
					screen,
					float32(lastX), float32(lastY),
					float32(minimapX), float32(minimapY),
					1, m.terrainColor, false,
				)
			}
		}

		lastX = float64(minimapX)
		lastY = float64(minimapY)
	}

	// Draw the viewport indicator (white rectangle)
	vector.StrokeLine(
		screen,
		float32(viewportX), float32(m.posY),
		float32(viewportX+viewportWidth), float32(m.posY),
		2, m.viewportColor, false,
	)
	vector.StrokeLine(
		screen,
		float32(viewportX), float32(m.posY+m.height),
		float32(viewportX+viewportWidth), float32(m.posY+m.height),
		2, m.viewportColor, false,
	)
	vector.StrokeLine(
		screen,
		float32(viewportX), float32(m.posY),
		float32(viewportX), float32(m.posY+m.height),
		2, m.viewportColor, false,
	)
	vector.StrokeLine(
		screen,
		float32(viewportX+viewportWidth), float32(m.posY),
		float32(viewportX+viewportWidth), float32(m.posY+m.height),
		2, m.viewportColor, false,
	)
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
