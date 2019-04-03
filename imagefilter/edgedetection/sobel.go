package edgedetection

import (
	"../convolution"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"time"
)

func SobelGradientCalculation(img draw.Image) (draw.Image, []float64) {
	start := time.Now()

	xKernel := []float64{
		-1, 0, 1,
		-2, 0, 2,
		-1, 0, 1}
	yKernel := []float64{
		1, 2, 1,
		0, 0, 0,
		-1, -2, -1}

	xGradient := convolution.ApplyConvolutionFilterRaw(img, xKernel, 1)
	yGradient := convolution.ApplyConvolutionFilterRaw(img, yKernel, 1)

	bounds := img.Bounds()
	gradient := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	unnormalizedGradient := make([]float64, bounds.Dx()*bounds.Dy())
	thetas := make([]float64, bounds.Dx()*bounds.Dy(), bounds.Dx()*bounds.Dy())
	max := 0.0

	for y := gradient.Bounds().Min.Y; y < gradient.Bounds().Max.Y; y++ {
		for x := gradient.Bounds().Min.X; x < gradient.Bounds().Max.X; x++ {
			index := y*bounds.Dx() + x

			xG := xGradient[index]
			yG := yGradient[index]
			unnormalizedGradient[index] = math.Sqrt(xG*xG + yG*yG)
			thetas[index] = math.Atan2(yG, xG)

			if max < unnormalizedGradient[index] {
				max = unnormalizedGradient[index]
			}
		}
	}

	// Normalize gradient
	for y := gradient.Bounds().Min.Y; y < gradient.Bounds().Max.Y; y++ {
		for x := gradient.Bounds().Min.X; x < gradient.Bounds().Max.X; x++ {
			index := y*bounds.Dx() + x
			value := uint8(unnormalizedGradient[index] / max * 255)

			gradient.Set(x, y, color.RGBA{R: value, G: value, B: value, A: 255})

		}
	}

	log.Println("Sobel:", time.Since(start))

	return gradient, thetas
}
