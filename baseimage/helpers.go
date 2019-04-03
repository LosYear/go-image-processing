package baseimage

import (
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
