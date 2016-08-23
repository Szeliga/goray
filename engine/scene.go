package engine

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Scene represents the s on which the scene is projected as a 2D picture
type Scene struct {
	Width, Height int
	Img           *image.RGBA
}

// NewScene returns a new Scene
func NewScene(width int, height int) *Scene {
	return &Scene{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// Iterate over the image s and call the provided function for each pixel
func (s *Scene) Iterate(colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.SetPixel(x, y, colorFunction(x, y))
		}
	}
}

// SetPixel sets the specified pixel to the specified color
func (s *Scene) SetPixel(x int, y int, color color.RGBA) {
	s.Img.Set(x, y, color)
}

// Save exports the image to hdd
func (s *Scene) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, s.Img)
}
