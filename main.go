package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/szeliga/goray/engine"
)

func main() {
	var width = 200
	var height = 200
	scene := engine.NewScene(width, height)
	scene.EachPixel(func(x, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / width),
			uint8(y * 255 / height),
			100,
			255,
		}
	})
	scene.Save(fmt.Sprintf("./renders/%d.png", time.Now().Unix()))
}
