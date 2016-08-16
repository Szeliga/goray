package engine

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"
)

// ImagePlane represents the plane on which the scene is projected as a 2D picture
type ImagePlane struct {
	Width, Height int
	Img           *image.RGBA
}

// NewImagePlane returns a new ImagePlane
func NewImagePlane(width int, height int) *ImagePlane {
	return &ImagePlane{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// SetPixel sets the specified pixel to the specified color
func (plane *ImagePlane) SetPixel(x int, y int, color color.RGBA) {
	plane.Img.Set(x, y, color)
}

// Render exports the image to hdd
func (plane *ImagePlane) Render() {
	f, err := os.Create(fmt.Sprintf("./renders/%d.png", time.Now().Unix()))

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f, plane.Img)
}

// Iterate over the image plane and call the provided function for each pixel
func (plane *ImagePlane) Iterate(pixelColor func(int, int) color.RGBA) {
	for x := 0; x < plane.Width; x++ {
		for y := 0; y < plane.Height; y++ {
			plane.SetPixel(x, y, pixelColor(x, y))
		}
	}
}
