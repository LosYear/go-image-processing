package edgedetection

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"time"
)

// TODO: if time is enough, make it concurrent
func NonMaxSuppression(img draw.Image, thetas []float64) draw.Image {
	start := time.Now()

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			index := y*bounds.Dx() + x

			angle := thetas[index] * 180 / math.Pi

			if angle < 0 {
				angle += 180
			}

			q := uint8(255)
			r := uint8(255)

			if (0 <= angle && angle < 22.5) || (157.5 <= angle && angle <= 180) {
				q = img.At(x+1, y).(color.RGBA).R
				r = img.At(x-1, y).(color.RGBA).R
			} else if 22.5 <= angle && angle < 67.5 {
				q = img.At(x-1, y+1).(color.RGBA).R
				r = img.At(x+1, y-1).(color.RGBA).R
			} else if 67.5 <= angle && angle < 112.5 {
				q = img.At(x, y+1).(color.RGBA).R
				r = img.At(x, y-1).(color.RGBA).R
			} else if 112.5 <= angle && angle < 157.5 {
				q = img.At(x-1, y-1).(color.RGBA).R
				r = img.At(x+1, y+1).(color.RGBA).R
			}

			imgPixel := img.At(x, y).(color.RGBA).R
			resultValue := uint8(0)

			if (imgPixel >= q) && (imgPixel >= r) {
				resultValue = imgPixel
			}

			result.Set(x, y, color.RGBA{R: resultValue, G: resultValue, B: resultValue, A: 255})
		}
	}

	log.Println("Non-maximum suppression:", time.Since(start))

	return result
}
