package main

import (
	"image/color"

	"github.com/szeliga/goray/engine"
)

func main() {
	plane := engine.NewImagePlane(200, 200)
	for x := 0; x < plane.Width; x++ {
		for y := 0; y < plane.Height; y++ {
			plane.SetPixel(x, y, color.RGBA{
				uint8(x * 255 / plane.Width),
				uint8(y * 255 / plane.Height),
				55,
				255,
			})
		}
	}
	plane.Render()
}
