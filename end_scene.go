package main

import "github.com/hajimehoshi/ebiten/v2"

type EndScene struct {
	sceneManager *SceneManager
}

func (e *EndScene) Update() error {
	return nil
}

func (e *EndScene) Draw(screen *ebiten.Image) {}

func (e *EndScene) Layout(outerWidth, outerHeight int) (int, int) {
	return outerWidth, outerHeight
}

func NewEndScene(sm *SceneManager) *EndScene {
	end := &EndScene{
		sceneManager: sm,
	}
	return end
}
