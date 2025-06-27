package assets

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed *
var assets embed.FS

var PlayerSprite = loadImage("sprites/player.png")
var AudioContext = audio.NewContext(44100)

var LaserSound = loadPlayerFromWav("audio/laser.wav", 0.3)




func loadImage(filePath string) *ebiten.Image {
	data, err := assets.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg
}

func loadPlayerFromWav(filePath string, volume float64) *audio.Player {
	data, err := assets.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	stream, err := wav.DecodeWithSampleRate(AudioContext.SampleRate(), bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	player, err := AudioContext.NewPlayer(stream)
	if err != nil {
		panic(err)
	}
	player.SetVolume(volume)

	return player
}
