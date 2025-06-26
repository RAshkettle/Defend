package main

import "github.com/hajimehoshi/ebiten/v2"

type GameScene struct {
	sceneManager *SceneManager
	player       *Player
}

func (g *GameScene) Update() error {
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	// Draw the player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.Image, op)
}

func (g *GameScene) Layout(outerWidth, outerHeight int) (int, int) {
	return 320, 240
}

func NewGameScene(sm *SceneManager) *GameScene {
	game := &GameScene{
		sceneManager: sm,
		player:       NewPlayer(),
	}
	return game
}
