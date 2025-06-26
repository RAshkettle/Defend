package main

type Camera struct {
	X      float64
	Speed  float64
	Width  float64
	Height float64
}

func NewCamera(screenWidth, screenHeight float64) *Camera {
	return &Camera{
		X:      0,
		Speed:  3.0,
		Width:  screenWidth,
		Height: screenHeight,
	}
}

func (c *Camera) MoveLeft(terrainWidth float64) {
	c.X -= c.Speed
	if c.X < 0 {
		c.X += terrainWidth
	}
}

func (c *Camera) MoveRight(terrainWidth float64) {
	c.X += c.Speed
	if c.X >= terrainWidth {
		c.X -= terrainWidth
	}
}
