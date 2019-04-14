package houghtransform

import (
	"./circle"
	"./line"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"math"
)

func DrawHoughLinesSet(img draw.Image, params []line.DetectedLineParams) draw.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Bounds().Dy()))
	draw.Draw(result, bounds, img, bounds.Min, draw.Src)

	gc := draw2dimg.NewGraphicContext(result)
	gc.SetFillColor(color.RGBA{255, 0, 0, 255})
	gc.SetStrokeColor(color.RGBA{255, 0, 0, 255})
	gc.SetLineWidth(3)

	gc.BeginPath()

	for _, p := range params {
		calc := line.ConvertToInterceptForm(p.Distance, (float64(p.Theta) / 180 * math.Pi))

		if p.Theta == 0 {
			gc.MoveTo(float64(p.Distance), float64(bounds.Min.Y))
			gc.LineTo(float64(p.Distance), float64(bounds.Max.X))

		} else {
			gc.MoveTo(0, -float64(calc(0)))
			gc.LineTo(float64(bounds.Max.X), -float64(calc(bounds.Max.X)))
		}
		gc.Stroke()
	}

	gc.Close()
	gc.FillStroke()

	return result
}

func DrawHoughCirclesSet(img draw.Image, params []circle.DetectedCircleParams) draw.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Bounds().Dy()))
	draw.Draw(result, bounds, img, bounds.Min, draw.Src)

	gc := draw2dimg.NewGraphicContext(result)
	gc.SetFillColor(color.RGBA{255, 0, 0, 255})
	gc.SetStrokeColor(color.RGBA{255, 0, 0, 255})
	gc.SetLineWidth(3)

	gc.BeginPath()

	for _, p := range params {

		gc.MoveTo(float64(p.A), float64(p.B))
		gc.LineTo(float64(p.A-p.Radius), float64(p.B))
		gc.Stroke()
	}

	gc.Close()
	gc.FillStroke()

	return result
}
