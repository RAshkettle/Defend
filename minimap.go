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
		terrainColor:  color.RGBA{255, 140, 0, 255},
		viewportColor: color.RGBA{255, 255, 255, 255},
		borderColor:   color.RGBA{0, 0, 255, 255},
	}
}

func (m *Minimap) Draw(screen *ebiten.Image, aliens []*Alien) {
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX+m.width), float32(m.posY), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY+m.height), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX+m.width), float32(m.posY), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)

	terrainWidth := m.terrain.width
	screenWidth := float64(screen.Bounds().Dx())
	maxHeight := 0.0
	for _, h := range m.terrain.points {
		if h > maxHeight {
			maxHeight = h
		}
	}
	scaleY := float64(m.height) / (maxHeight * 1.2)
	cameraX := m.camera.X
	for cameraX < 0 {
		cameraX += terrainWidth
	}
	for cameraX >= terrainWidth {
		cameraX -= terrainWidth
	}
	viewportWidth := int((screenWidth / terrainWidth) * float64(m.width))
	viewportPosition := float64(cameraX) / terrainWidth
	viewportX := m.posX + int(float64(m.width)*viewportPosition) - viewportWidth/2
	if viewportX < m.posX {
		viewportX += m.width
	}
	if viewportX+viewportWidth > m.posX+m.width {
		viewportX -= m.width
	}
	pointSpacing := 5.0
	minimapScale := float64(m.width) / terrainWidth
	for minimapX := 0; minimapX < m.width-1; minimapX++ {
		worldX := float64(minimapX) / minimapScale
		index := int(worldX/pointSpacing) % len(m.terrain.points)
		nextIndex := (index + 1) % len(m.terrain.points)
		if index >= 0 && index < len(m.terrain.points) {
			height := m.terrain.points[index]
			nextHeight := m.terrain.points[nextIndex]
			y1 := m.posY + m.height - int(height*scaleY)
			y2 := m.posY + m.height - int(nextHeight*scaleY)
			if y1 < m.posY {
				y1 = m.posY
			}
			if y1 > m.posY+m.height {
				y1 = m.posY + m.height
			}
			if y2 < m.posY {
				y2 = m.posY
			}
			if y2 > m.posY+m.height {
				y2 = m.posY + m.height
			}
			vector.StrokeLine(
				screen,
				float32(m.posX+minimapX), float32(y1),
				float32(m.posX+minimapX+1), float32(y2),
				1, m.terrainColor, false,
			)
		}
	}
	alienColor := color.RGBA{255, 255, 255, 255}
	for _, alien := range aliens {
		if alien.Active {
			alienWorldX := alien.X
			for alienWorldX < 0 {
				alienWorldX += terrainWidth
			}
			for alienWorldX >= terrainWidth {
				alienWorldX -= terrainWidth
			}
			alienPosition := float64(alienWorldX) / terrainWidth
			alienMinimapX := m.posX + int(float64(m.width)*alienPosition)
			if alienMinimapX < m.posX {
				alienMinimapX += m.width
			}
			if alienMinimapX >= m.posX+m.width {
				alienMinimapX -= m.width
			}
			alienMinimapY := m.posY + m.height - int(alien.Y*scaleY) - 2
			if alienMinimapY < m.posY {
				alienMinimapY = m.posY
			}
			if alienMinimapY > m.posY+m.height {
				alienMinimapY = m.posY + m.height
			}
			vector.StrokeLine(
				screen,
				float32(alienMinimapX),
				float32(alienMinimapY),
				float32(alienMinimapX+1),
				float32(alienMinimapY),
				1,
				alienColor,
				false,
			)
		}
	}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
