package convolution

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"runtime"
	"sync"
)

func ApplyConvolutionKernel(pixels []uint8, kernel []float64, c float64) float64 {
	if len(pixels) != len(kernel) {
		log.Fatal("Sizes of kernel and data slice have to be equal")
	}

	result := 0.0
	l := len(kernel) - 1
	for i := 0; i < len(pixels); i++ {
		result += kernel[l-i] * float64(pixels[i])
	}

	return result * c
}

func applyConvolutionFilterRoutine(wg *sync.WaitGroup, img draw.Image, result draw.Image, kernel []float64, c float64, from, to int) {
	defer wg.Done()
	bounds := img.Bounds()
	sliceSize := uint(math.Sqrt(float64(len(kernel))))

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixels := SlicePixels(img, x, y, sliceSize, 0)
			convolutedPixel := uint8(ApplyConvolutionKernel(pixels, kernel, c))

			result.Set(x, y, color.RGBA{R: convolutedPixel, G: convolutedPixel, B: convolutedPixel, A: 255})
		}
	}

}

func applyConvolutionFilterRawRoutine(wg *sync.WaitGroup, img draw.Image, result []float64, kernel []float64, c float64, from, to int) {
	defer wg.Done()
	bounds := img.Bounds()
	sliceSize := uint(math.Sqrt(float64(len(kernel))))

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixels := SlicePixels(img, x, y, sliceSize, 0)
			convolutedPixel := ApplyConvolutionKernel(pixels, kernel, c)

			index := y*bounds.Dx() + x
			result[index] = convolutedPixel
		}
	}

}

func ApplyConvolutionFilterRaw(img draw.Image, kernel []float64, c float64) []float64 {
	bounds := img.Bounds()
	result := make([]float64, bounds.Dy()*bounds.Dx(), bounds.Dy()*bounds.Dx())

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

		go applyConvolutionFilterRawRoutine(&wg, img, result, kernel, c, start, end)
	}

	wg.Wait()
	return result
}

func ApplyConvolutionFilter(img draw.Image, kernel []float64, c float64) draw.Image {
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

		go applyConvolutionFilterRoutine(&wg, img, result, kernel, c, start, end)
	}

	wg.Wait()
	return result
}
