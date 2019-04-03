package edgedetection

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
)

func calculateThresholds(img draw.Image, lowThresholdRatio, highThresholdRatio float64) (float64, float64) {
	highThreshold := img.At(0, 0).(color.RGBA).R

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA).R

			if pixel > highThreshold {
				highThreshold = pixel
			}
		}
	}

	return float64(highThreshold) * lowThresholdRatio, float64(highThreshold) * highThresholdRatio
}

func DoubleThreshold(img draw.Image, lowThresholdRatio, highThresholdRatio float64) (draw.Image, uint8, uint8) {
	start := time.Now()

	lowThreshold, highThreshold := calculateThresholds(img, lowThresholdRatio, highThresholdRatio)

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	weakColor, strongColor := color.RGBA{R: 25, G: 25, B: 25, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := float64(img.At(x, y).(color.RGBA).R)

			if pixel >= highThreshold {
				result.Set(x, y, strongColor)
			} else if pixel >= lowThreshold && pixel < highThreshold {
				result.Set(x, y, weakColor)

			} else {
				result.Set(x, y, color.RGBA{R: 0, B: 0, G: 0, A: 255})
			}

		}
	}

	log.Println("Double threshold:", time.Since(start))

	return result, weakColor.R, strongColor.R
}
