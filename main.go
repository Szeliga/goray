package main

import (
	"image/color"

	"github.com/szeliga/goray/engine"
)

func main() {
	plane := engine.NewImagePlane(200, 200)
	plane.Iterate(func(x int, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / plane.Width),
			uint8(y * 255 / plane.Height),
			100,
			255,
		}
	})
	plane.Render()
}
