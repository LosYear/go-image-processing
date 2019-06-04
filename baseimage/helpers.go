package baseimage

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

func GetPixelAtChannel(img draw.Image, x int, y int, c uint) uint8 {
	pixel := img.At(x, y).(color.RGBA)

	if c == 0 {
		return pixel.R
	} else if c == 1 {
		return pixel.G
	} else if c == 2 {
		return pixel.B
	} else {
		log.Fatal("Channel should be in [0, 1, 2]")
	}

	return 0
}

func SliceGrayscaleRegion(img draw.Image, startX, endX, startY, endY int) [][]uint8 {
	result := make([][]uint8, endY-startY, endY-startY)

	for row := range result {
		result[row] = make([]uint8, endX-startX, endX-startX)
	}

	for row := startY; row < endY; row++ {
		for column := startX; column < endX; column++ {
			result[row-startY][column-startX] = img.At(column, row).(color.RGBA).R
		}
	}

	return result
}

func FillRegionWithColor(img draw.Image, color color.RGBA, startX, endX, startY, endY int) draw.Image {

	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Bounds().Dy()))
	draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			result.Set(x, y, color)
		}
	}

	return result
}
