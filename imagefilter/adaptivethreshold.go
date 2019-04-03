package imagefilter

import (
	"./convolution"
	"image"
	"image/color"
	"image/draw"
	"log"
	"runtime"
	"sync"
	"time"
)

func adaptiveThresholdRoutine(wg *sync.WaitGroup, img draw.Image, result draw.Image, sliceSize uint, coeff int, from int, to int) {
	defer wg.Done()
	bounds := img.Bounds()
	sliceCenter := sliceSize * sliceSize / 2

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixels := convolution.SlicePixels(img, x, y, sliceSize, 0)
			pixel := pixels[sliceCenter]

			sum := 0

			for _, val := range pixels {
				sum += int(val)
			}

			threshold := float64(sum)/float64(len(pixels)) - float64(coeff)

			var newValue uint8

			if float64(pixel) > threshold {
				newValue = 255
			} else {
				newValue = 0
			}

			result.Set(x, y, color.RGBA{R: newValue, G: newValue, B: newValue, A: 255})
		}
	}
}

func AdaptiveThreshold(img draw.Image, sliceSize uint, coeff int) draw.Image {
	start := time.Now()

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Bounds().Dy()))

	threads := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup

	rowsPerThread := bounds.Dy() / threads

	for i := 0; i < threads; i++ {
		wg.Add(1)

		start := img.Bounds().Min.Y + rowsPerThread*i

		var end int

		if i == threads-1 {
			end = img.Bounds().Max.Y

		} else {
			end = start + rowsPerThread

		}

		go adaptiveThresholdRoutine(&wg, img, result, sliceSize, coeff, start, end)
	}

	wg.Wait()

	log.Println("Adaptive threshold:", time.Since(start))
	return result
}
