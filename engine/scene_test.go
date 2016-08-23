package engine

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSceneReturnsANewScene(t *testing.T) {
	scene := NewScene(4, 4)
	rect := image.Rect(0, 0, 4, 4)
	assert.Equal(t, 4, scene.Width, "sets width of the scene")
	assert.Equal(t, 4, scene.Height, "sets height of the scene")
	assert.True(t, assert.ObjectsAreEqualValues(rect, scene.Img.Bounds()),
		"creates an image.RGBA with proper bounds")
}

func TestSceneIterateSetsEachPixelToTheProvidedFunctionReturn(t *testing.T) {
	scene := NewScene(4, 4)
	c := randomColor()

	scene.Iterate(func(x, y int) color.RGBA { return c })
	img := generateImage(4, 4, c)
	assert.True(t, assert.ObjectsAreEqualValues(img, scene.Img),
		"sets every pixel of the image to the provided values")
}

func TestSceneSetPixelSetsTheSpecifiedPixelToTheSpecifiedColor(t *testing.T) {
	scene := NewScene(4, 4)
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	scene.SetPixel(0, 0, black)
	scene.SetPixel(3, 3, white)
	assert.True(t, assert.ObjectsAreEqualValues(scene.Img.At(0, 0), black),
		"sets the top left pixel to a black color")
	assert.True(t, assert.ObjectsAreEqualValues(scene.Img.At(3, 3), white),
		"sets the bottom right pixel to a white color")
}

func TestSceneSavePersistsTheFileOnTheDisk(t *testing.T) {
	scene := NewScene(4, 4)
	c := randomColor()
	scene.Iterate(func(x, y int) color.RGBA { return c })
	scene.Save("./temp.png")

	infile, err := os.Open("./temp.png")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	readImg, _, err := image.Decode(infile)
	if err != nil {
		log.Fatal(err)
	}

	assert.True(t, assert.ObjectsAreEqualValues(scene.Img, readImg),
		"saves the image to the disk")
	os.Remove("./temp.png")
}

func TestSceneSavePanicsWhenFileCannotBePersisted(t *testing.T) {
	scene := NewScene(4, 4)
	c := randomColor()
	scene.Iterate(func(x, y int) color.RGBA { return c })
	assert.Panics(t, func() {
		scene.Save("/etc/temp.png")
	}, "panics when the file cannot be persisted")
}

// Helper functions
func generateImage(w, h int, pixelColor color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			img.Set(x, y, pixelColor)
		}
	}
	return img
}

func randomColor() color.RGBA {
	rand := rand.New(rand.NewSource(99))
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}
