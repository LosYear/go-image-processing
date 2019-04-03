package imagefilter

import (
	"../../imagefilter"
	"image"
	"image/color"
	"image/draw"
	"log"
	"testing"
)

func generateImageFromPixelSet(data [][]uint8) draw.Image {
	result := image.NewRGBA(image.Rect(0, 0, len(data), len(data[0])))

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			pixel := uint8(data[i][j])
			result.Set(j, i, color.RGBA{R: pixel, G: pixel, B: pixel, A: 255})
		}
	}

	return result
}

func printImage(img draw.Image) {
	bounds := img.Bounds()
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			pixel := img.At(i, j).(color.RGBA).R
			log.Print(pixel)
		}
		log.Println()
	}
}

func TestSobel(t *testing.T) {
	data := [][]uint8{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	img := generateImageFromPixelSet(data)
	_, angles := imagefilter.SobelGradientCalculation(img)

	log.Println(angles)

}
