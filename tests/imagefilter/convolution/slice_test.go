package convolution

import (
	"../../../imagefilter/convolution"
	"image"
	"image/color"
	"image/draw"
	"reflect"
	"testing"
)

func fixture() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			value := uint8(y*10 + x)
			img.Set(x, y, color.RGBA{R: value, G: value, B: value, A: 255})
		}
	}

	return img
}

func genericSliceTest(t *testing.T, x int, y int, size uint, expected []uint8) {
	img := fixture()
	pixels := convolution.SlicePixels(img, x, y, size, 0)

	if !reflect.DeepEqual(pixels, expected) {
		t.Error("Got:", pixels,
			"expected:", expected)
	}
}

func TestSliceLeftBorder(t *testing.T) {
	genericSliceTest(t, 0, 0, 3, []uint8{0, 0, 1, 0, 0, 1, 10, 10, 11})
}

func TestSliceOnePixel(t *testing.T) {
	genericSliceTest(t, 0, 0, 1, []uint8{0})
}

func TestSliceMiddle(t *testing.T) {
	genericSliceTest(t, 3, 3, 3, []uint8{22, 23, 24, 32, 33, 34, 42, 43, 44})

}
