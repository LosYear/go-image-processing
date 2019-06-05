package rotate

import (
	"../baseimage"
	"image/color"
	"image/draw"
)

func Rotate(img draw.Image, startX, endX, startY, endY, angle int) draw.Image {
	data := baseimage.SliceGrayscaleRegion(img, startX, endX, startY, endY)
	img = baseimage.FillRegionWithColor(img, color.RGBA{255, 255, 255, 255}, startX, endX, startY, endY)

	centerX := startX + (endX-startX)/2
	centerY := startY + (endY-startY)/2

	bounds := img.Bounds()

	for row := bounds.Min.Y; row < bounds.Max.Y; row++ {
		for column := bounds.Min.X; column < bounds.Max.X; column++ {
			sourceX, sourceY := mapDestinationToSource(column, row, centerX, centerY, angle)

			if sourceX < 0 ||
				sourceX >= bounds.Max.X ||
				sourceX < startX ||
				sourceX >= endX ||
				sourceY < 0 ||
				sourceY >= bounds.Max.Y ||
				sourceY < startY ||
				sourceY >= endY {
				continue
			}

			sourceX = sourceX - startX
			sourceY = sourceY - startY

			pixel := data[sourceY][sourceX]
			clr := color.RGBA{R: pixel, G: pixel, B: pixel, A: 255}

			img.Set(column, row, clr)
		}
	}

	return img
}
