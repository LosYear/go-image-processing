package model

import (
	"math"
)

func rotatePoint(x, y, centerX, centerY, angle float64) (float64, float64) {
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	x -= centerX
	y -= centerY

	newX := x*cos - y*sin
	newY := x*sin + y*cos

	return newX + centerX, newY + centerY

}

func RotateModel(model VectorModel, angle int) VectorModel {
	var result VectorModel

	dimensions := result.Dimensions
	modelWidth, modelHeight := dimensions.width, dimensions.height
	radiansAngle := float64(angle) * math.Pi / 180

	for _, polygon := range model.Polygons {
		var newPolygon VectorPolygon
		newPolygon.Transparent = polygon.Transparent

		for _, point := range polygon.Points {
			newX, newY := rotatePoint(point.X, point.Y, modelWidth/2, modelHeight/2, radiansAngle)

			newPolygon.Points = append(newPolygon.Points, Point{X: newX, Y: newY})
		}

		result.Polygons = append(result.Polygons, newPolygon)
	}

	result.Dimensions = FillModelDimensions(result)

	return result
}
