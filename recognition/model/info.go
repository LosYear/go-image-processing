package model

import "math"

func FillModelDimensions(model VectorModel) Dimensions {
	dimensions := Dimensions{width: 0, height: 0, minX: model.Polygons[0].Points[0].X, minY: model.Polygons[0].Points[0].Y}

	modelWidth := 0.0
	modelHeight := 0.0

	for _, polygon := range model.Polygons {
		for _, point := range polygon.Points {
			modelWidth = math.Max(modelWidth, point.X)
			modelHeight = math.Max(modelHeight, point.Y)

			dimensions.minX = math.Min(dimensions.minX, point.X)
			dimensions.minY = math.Min(dimensions.minY, point.Y)
		}
	}

	dimensions.width = modelWidth - dimensions.minX
	dimensions.height = modelHeight - dimensions.minY

	return dimensions
}
