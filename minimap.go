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
	scaleY := float64(m.height) / m.camera.Height // Match main view's Y scaling

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

	// Position the viewport indicator based on the camera's left edge
	viewportPosition := float64(cameraX) / terrainWidth
	viewportX := m.posX + int(float64(m.width)*viewportPosition)

	// Draw all terrain points using same pointSpacing as in terrain.Draw
	pointSpacing := 5.0
	lastX := -1.0
	lastY := -1.0

	// Draw all terrain points for all 3 screens
	for i := 0; i < len(m.terrain.points); i++ {
		// Convert world position to minimap position
		worldX := float64(i) * pointSpacing
		minimapX := m.posX + int(worldX*scaleX)

		// Calculate height at this point
		height := m.terrain.points[i]
		minimapY := m.posY + m.height - int(height*scaleY)

		// Draw line segment if not the first point and not wrapping around edge
		if lastX >= 0 {
			vector.StrokeLine(
				screen,
				float32(lastX), float32(lastY),
				float32(minimapX), float32(minimapY),
				1, m.terrainColor, false,
			)
		}

		lastX = float64(minimapX)
		lastY = float64(minimapY)
	}

	// Draw the viewport indicator, handling wrapping by drawing two boxes if needed.
	if viewportX+viewportWidth > m.posX+m.width {
		// --- View is wrapped, draw two boxes ---
		// 1. Draw the part on the right edge of the minimap
		part1Width := (m.posX + m.width) - viewportX
		vector.StrokeRect(screen, float32(viewportX), float32(m.posY), float32(part1Width), float32(m.height), 2, m.viewportColor, false)

		// 2. Draw the part on the left edge of the minimap
		part2Width := (viewportX + viewportWidth) - (m.posX + m.width)
		vector.StrokeRect(screen, float32(m.posX), float32(m.posY), float32(part2Width), float32(m.height), 2, m.viewportColor, false)
	} else {
		// --- View is not wrapped, draw a single box ---
		vector.StrokeRect(screen, float32(viewportX), float32(m.posY), float32(viewportWidth), float32(m.height), 2, m.viewportColor, false)
	}
}
