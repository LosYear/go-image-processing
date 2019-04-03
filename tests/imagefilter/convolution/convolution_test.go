package convolution

import (
	"../../../imagefilter/convolution"
	"testing"
)

func genericConvolutionKernelTest(t *testing.T, pixels []uint8, kernel []float64, expected float64) {
	result := convolution.ApplyConvolutionKernel(pixels, kernel, 1)

	if expected != result {
		t.Error("Got:", pixels,
			"expected:", expected)
	}
}

func TestConvolutionKernel1(t *testing.T) {
	data := []uint8{12, 14, 41, 43, 84, 24, 2, 1, 43}
	kernel := []float64{0.5, 0.75, 0.5, 0.75, 1.0, 0.75, 0.5, 0.75, 0.5}

	genericConvolutionKernelTest(t, data, kernel, 194.5)
}

func TestConvolutionKernel2(t *testing.T) {
	data := []uint8{185, 185, 183, 143, 155, 90, 153, 147, 161}
	kernel := []float64{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1}

	genericConvolutionKernelTest(t, data, kernel, 140.2)
}
