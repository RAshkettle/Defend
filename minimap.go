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
	minimapHeight := 30
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

func (m *Minimap) drawBorder(screen *ebiten.Image) {
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX+m.width), float32(m.posY), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY+m.height), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX), float32(m.posY), float32(m.posX), float32(m.posY+m.height), 1, m.borderColor, false)
	vector.StrokeLine(screen, float32(m.posX+m.width), float32(m.posY), float32(m.posX+m.width), float32(m.posY+m.height), 1, m.borderColor, false)
}

func (m *Minimap) drawTerrain(screen *ebiten.Image) {
	terrainWidth := m.terrain.width
	scaleX := float64(m.width) / terrainWidth

	maxHeight := 0.0
	for _, h := range m.terrain.points {
		if h > maxHeight {
			maxHeight = h
		}
	}

	scaleY := scaleX

	yOffset := float64(m.posY) + float64(m.height)*0.9 - maxHeight*scaleY/2

	cameraX := m.camera.X
	for cameraX < 0 {
		cameraX += terrainWidth
	}
	for cameraX >= terrainWidth {
		cameraX -= terrainWidth
	}

	pointSpacing := 5.0

	for worldX := 0.0; worldX < terrainWidth; worldX += pointSpacing {
		screenX := (worldX * scaleX)

		index := int(worldX / pointSpacing)
		nextIndex := (index + 1) % len(m.terrain.points)

		if index >= 0 && index < len(m.terrain.points) {
			height := m.terrain.points[index]
			nextHeight := m.terrain.points[nextIndex]

			y1 := yOffset - height*scaleY
			y2 := yOffset - nextHeight*scaleY
			x1 := float64(m.posX) + screenX
			x2 := float64(m.posX) + screenX + (pointSpacing * scaleX)

			vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, m.terrainColor, false)
		}
	}
}

func (m *Minimap) drawViewport(screen *ebiten.Image) {
	terrainWidth := m.terrain.width
	scaleX := float64(m.width) / terrainWidth

	cameraX := m.camera.X
	for cameraX < 0 {
		cameraX += terrainWidth
	}
	for cameraX >= terrainWidth {
		cameraX -= terrainWidth
	}

	screenWidth := m.camera.Width
	viewportWidth := int((screenWidth / terrainWidth) * float64(m.width))

	viewportPosition := float64(m.posX) + cameraX*scaleX

	viewportX := int(viewportPosition)

	if viewportX < m.posX {
		viewportX = m.posX
	}
	if viewportX+viewportWidth > m.posX+m.width {
		if cameraX > terrainWidth-screenWidth {
			viewportX = m.posX + int((terrainWidth-screenWidth)*scaleX)
		} else {
			viewportX = m.posX
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

func (m *Minimap) Draw(screen *ebiten.Image) {
	m.drawBorder(screen)
	m.drawTerrain(screen)
	m.drawViewport(screen)
}


