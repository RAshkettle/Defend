package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	sceneManager *SceneManager
	player       *Player
	camera       *Camera
	terrain      *Terrain
	minimap      *Minimap
}

func (g *GameScene) Update() error {
	if err := g.player.Update(g.camera, float64(g.terrain.width)); err != nil {
		return err
	}

	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.terrain.Draw(screen, g.camera)

	op := &ebiten.DrawImageOptions{}

	if g.player.Facing == LEFT {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(g.player.Image.Bounds().Dx()), 0)
	}

	op.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(g.player.Image, op)

	// Draw the minimap overlay
	g.minimap.Draw(screen)
}

func (g *GameScene) Layout(outerWidth, outerHeight int) (int, int) {
	return 320, 240
}

func NewGameScene(sm *SceneManager) *GameScene {
	width, height := 320.0, 240.0
	camera := NewCamera(width, height)
	terrain := NewTerrain(width)

	game := &GameScene{
		sceneManager: sm,
		player:       NewPlayer(),
		camera:       camera,
		terrain:      terrain,
	}

	// Initialize minimap after terrain and camera are created
	game.minimap = NewMinimap(terrain, camera, int(width), int(height))

	return game
}
