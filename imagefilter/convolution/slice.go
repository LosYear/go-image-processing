package convolution

import (
	"../../baseimage"
	"image/draw"
)

type offset struct {
	x int
	y int
}

func SlicePixels(img draw.Image, x int, y int, size uint, c uint) []uint8 {
	pixels := make([]uint8, 0, size*size)
	bounds := img.Bounds()

	for offset := range generateOffsets(size) {
		xOffseted := x + offset.x
		yOffseted := y + offset.y

		if xOffseted < bounds.Min.X {
			xOffseted = bounds.Min.X
		} else if xOffseted >= bounds.Max.X {
			xOffseted = bounds.Max.X - 1
		}

		if yOffseted < bounds.Min.Y {
			yOffseted = bounds.Min.Y
		} else if yOffseted >= bounds.Max.Y {
			yOffseted = bounds.Max.Y - 1
		}

		pixels = append(pixels, baseimage.GetPixelAtChannel(img, xOffseted, yOffseted, c))

	}

	return pixels
}

func generateOffsets(size uint) chan offset {
	ch := make(chan offset)

	absOffset := int(size / 2)

	go func() {
		for i := -1 * absOffset; i <= absOffset; i++ {
			for j := -1 * absOffset; j <= absOffset; j++ {
				ch <- offset{x: j, y: i}
			}

		}

		close(ch)
	}()

	return ch
}
