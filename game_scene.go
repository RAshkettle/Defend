package main

import (
	"defend/assets"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	sceneManager *SceneManager
	player       *Player
	camera       *Camera
	terrain      *Terrain
	minimap      *Minimap
	aliens       []*Alien
}

func (g *GameScene) Update() error {
	if err := g.player.Update(g.camera, float64(g.terrain.width)); err != nil {
		return err
	}

	for _, a := range g.aliens {
		a.Update(g.terrain.width)
	}
	for _, laser := range g.player.ActiveShots {
		laser.CheckAlienCollision(g.aliens, g.terrain.width)
	}
	// Clean up any dead aliens
	activeAliens := make([]*Alien, 0)
	for _, a := range g.aliens {
		if a.Active {
			activeAliens = append(activeAliens, a)
		}
	}
	g.aliens = activeAliens

	g.aliens = CheckAlienSpawn(g.aliens, g.terrain.width)
	g.CheckPlayerAlienCollision()
	return nil
}

func (g *GameScene) CheckPlayerAlienCollision() {
	playerWorldX := g.player.X + g.camera.X
	p := image.Rect(int(playerWorldX), int(g.player.Y), int(playerWorldX)+g.player.Image.Bounds().Dx(), int(g.player.Y)+g.player.Image.Bounds().Dy())

	for _, a := range g.aliens {
		// Check collision at alien's current position
		alienRect := image.Rect(int(a.X), int(a.Y), int(a.X)+a.Image.Bounds().Dx(), int(a.Y)+a.Image.Bounds().Dy())
		if alienRect.Overlaps(p) {
			a.Active = false
			assets.PlayerDeathSound.Rewind()
			assets.PlayerDeathSound.Play()
			g.sceneManager.TransitionTo(SceneEndScreen)
			return
		}

		// Check collision with alien wrapped to the left
		wrappedAlienRectLeft := image.Rect(int(a.X-g.terrain.width), int(a.Y), int(a.X-g.terrain.width)+a.Image.Bounds().Dx(), int(a.Y)+a.Image.Bounds().Dy())
		if wrappedAlienRectLeft.Overlaps(p) {
			a.Active = false
			assets.PlayerDeathSound.Rewind()
			assets.PlayerDeathSound.Play()
			g.sceneManager.TransitionTo(SceneEndScreen)
			return
		}

		// Check collision with alien wrapped to the right
		wrappedAlienRectRight := image.Rect(int(a.X+g.terrain.width), int(a.Y), int(a.X+g.terrain.width)+a.Image.Bounds().Dx(), int(a.Y)+a.Image.Bounds().Dy())
		if wrappedAlienRectRight.Overlaps(p) {
			a.Active = false
			assets.PlayerDeathSound.Rewind()
			assets.PlayerDeathSound.Play()
			g.sceneManager.TransitionTo(SceneEndScreen)
			return
		}
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.terrain.Draw(screen, g.camera)
	for _, alien := range g.aliens {
		alien.Draw(screen, g.camera, g.terrain.width)
	}
	for _, laser := range g.player.ActiveShots {
		laser.Draw(screen, g.camera)
	}
	op := &ebiten.DrawImageOptions{}
	if g.player.Facing == LEFT {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(g.player.Image.Bounds().Dx()), 0)
	}
	op.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.Image, op)
	g.minimap.Draw(screen, g.aliens)
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
		aliens:       []*Alien{},
	}
	game.minimap = NewMinimap(terrain, camera, int(width), int(height))
	return game
}
