package imagefilter

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
)

const rC = 0.299
const gC = 0.587
const bC = 0.114

func Grayscale(img draw.Image) draw.Image {
	start := time.Now()

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Bounds().Dy()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA)

			grayscaleValue := uint8(float64(pixel.R)*rC + float64(pixel.G)*gC + float64(pixel.B)*bC)

			result.Set(x, y, color.RGBA{R: grayscaleValue, G: grayscaleValue, B: grayscaleValue, A: 255})
		}
	}

	log.Println("Grayscale:", time.Since(start))

	return result
}
