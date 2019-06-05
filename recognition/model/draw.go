package model

import (
	"../../imagefilter"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
)

func DrawModel(model VectorModel, width, height int) draw.Image {
	result := image.NewRGBA(image.Rect(0, 0, width, height))

	for row := result.Bounds().Min.Y; row < result.Bounds().Max.Y; row++ {
		for column := result.Bounds().Min.X; column < result.Bounds().Max.X; column++ {
			result.SetRGBA(column, row, color.RGBA{R: 255, B: 255, G: 255, A: 255})
		}
	}

	gc := draw2dimg.NewGraphicContext(result)
	gc.SetLineWidth(1)

	dimensions := model.Dimensions
	modelWidth, modelHeight := dimensions.width, dimensions.height

	for _, polygon := range model.Polygons {
		gc.BeginPath()
		for index, point := range polygon.Points {
			scaledX := (point.X - dimensions.minX) / modelWidth * float64(width)
			scaledY := (point.Y - dimensions.minY) / modelHeight * float64(height)

			if index == 0 {
				gc.MoveTo(scaledX, scaledY)
			} else {
				gc.LineTo(scaledX, scaledY)
			}
		}
		if polygon.Transparent {
			gc.SetFillColor(color.RGBA{255, 255, 255, 255})
			gc.SetStrokeColor(color.RGBA{255, 255, 255, 255})
		} else {
			gc.SetFillColor(color.RGBA{0, 0, 0, 255})
			gc.SetStrokeColor(color.RGBA{0, 0, 0, 255})
		}
		gc.FillStroke()
		gc.Close()
	}
	return imagefilter.Threshold(result, 100)
}
