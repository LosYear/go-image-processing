package edgedetection

import (
	"../convolution"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
)

func maxInSlice(slice []uint8) uint8 {
	max := slice[0]

	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func Hysterisis(img draw.Image, weakColor, strongColor uint8) draw.Image {
	start := time.Now()

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA).R

			if pixel == weakColor {
				pixels := convolution.SlicePixels(img, x, y, 3, 0)
				max := maxInSlice(pixels)

				if max >= strongColor {
					result.Set(x, y, color.RGBA{R: strongColor, G: strongColor, B: strongColor, A: 255})
				} else {
					result.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 255})
				}
			} else {
				result.Set(x, y, color.RGBA{R: pixel, G: pixel, B: pixel, A: 255})
			}

		}
	}

	log.Println("Hysterisis:", time.Since(start))

	return result
}
