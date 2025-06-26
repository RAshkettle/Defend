package main

import "github.com/hajimehoshi/ebiten/v2"

type TitleScene struct {
	sceneManager *SceneManager
}

func (t *TitleScene) Update() error {
	return nil
}

func (t *TitleScene) Draw(screen *ebiten.Image) {}

func (t *TitleScene) Layout(outerWidth, outerHeight int) (int, int) {
	return outerWidth, outerHeight
}

func NewTitleScene(sm *SceneManager) *TitleScene {
	title := &TitleScene{
		sceneManager: sm,
	}
	return title
}
