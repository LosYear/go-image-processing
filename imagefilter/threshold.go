package imagefilter

import (
	"image"
	"image/color"
	"image/draw"
	"runtime"
	"sync"
)

func thresholdRoutine(wg *sync.WaitGroup, img draw.Image, result draw.Image, threshold, from int, to int) {
	defer wg.Done()
	bounds := img.Bounds()

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA)
			var newValue uint8

			if int(pixel.R) > threshold {
				newValue = 255
			} else {
				newValue = 0
			}

			result.Set(x, y, color.RGBA{R: newValue, G: newValue, B: newValue, A: 255})
		}
	}
}

func Threshold(img draw.Image, threshold int) draw.Image {
	//start := time.Now()

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

		go thresholdRoutine(&wg, img, result, threshold, start, end)
	}

	wg.Wait()

	//log.Println("Threshold:", time.Since(start))
	return result
}
