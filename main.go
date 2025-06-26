package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	sm := NewSceneManager()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Defend")
	ebiten.SetWindowSize(640, 480)

	err := ebiten.RunGame(sm)
	if err != nil {
		panic(err)
	}
}
