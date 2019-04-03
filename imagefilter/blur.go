package imagefilter

import (
	"image/draw"
	"log"
	"time"
)
import "./convolution"

//func generateGaussianKernel(kernelSize uint, sigma float64) []uint8 {
//	kernel := make([]uint8, kernelSize*kernelSize, kernelSize*kernelSize)
//	kernelSide := int(kernelSize)
//	sum := 0
//
//	for i := 1; i <= kernelSide; i++ {
//		for j := 1; j <= kernelSide; j++ {
//			index := (i-1)*kernelSide + (j - 1)
//
//			expVal := (math.Pow(float64(i-(kernelSide+1)), 2) + math.Pow(float64(j-(kernelSide+1)), 2)) / (2 * sigma * sigma)
//			value := uint8(1 / (2.0 * math.Pi * sigma * sigma) * math.Exp(-1*expVal))
//
//			log.Println(1 / (2.0 * math.Pi * sigma * sigma) * math.Exp(-1*expVal) * 100)
//
//			sum += int(value)
//
//			kernel[index] = value
//		}
//	}
//	return kernel
//}

func GaussianBlur(img draw.Image) draw.Image {
	start := time.Now()

	kernel := []float64{
		1, 4, 7, 4, 1,
		4, 16, 26, 16, 4,
		7, 26, 41, 26, 7,
		4, 16, 26, 16, 4,
		1, 4, 7, 4, 1}

	result := convolution.ApplyConvolutionFilter(img, kernel, 1.0/273.0)
	log.Println("Blur:", time.Since(start))

	return result
}
