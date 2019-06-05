package recognition

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)
import "./model"

func matchPercentage(img draw.Image, model draw.Image) float64 {
	totalDots := 0
	matchedDots := 0

	bounds := model.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			modelPixel := model.At(x, y).(color.RGBA)
			imagePixel := img.At(x, y).(color.RGBA)

			if (modelPixel.R == 0 && modelPixel.G == 0 && modelPixel.B == 0 && modelPixel.A > 0) ||
				(imagePixel.R == 0 && imagePixel.G == 0 && imagePixel.B == 0 && imagePixel.A > 0) {
				totalDots++
			}

			if (modelPixel.R == 0 && modelPixel.G == 0 && modelPixel.B == 0 && modelPixel.A > 0) &&
				(imagePixel.R == 0 && imagePixel.G == 0 && imagePixel.B == 0 && imagePixel.A > 0) {
				matchedDots++
			}
		}
	}

	return float64(matchedDots) / float64(totalDots)
}

func recognizeEnitity(img draw.Image, models []model.VectorModel) []float64 {
	probabilities := make([]float64, len(models), len(models))

	for mIdx, m := range models {
		//m := models[2]
		maxProb := 0.0

		for angle := -30; angle <= 30; angle++ {
			rotatedModel := model.RotateModel(m, angle)
			renderedModel := model.DrawModel(rotatedModel, img.Bounds().Dx(), img.Bounds().Dy())

			maxProb = math.Max(maxProb, matchPercentage(img, renderedModel))
		}
		probabilities[mIdx] = maxProb
	}

	return probabilities
}

func Recognize(img draw.Image, models []model.VectorModel) []int {
	boxes := FindBoundingBoxes(img)

	regions := make([]draw.Image, len(boxes), len(boxes))

	// Draw bounds
	for i, box := range boxes {
		region := image.NewRGBA(image.Rect(0, 0, box.MaxX-box.MinX+1, box.MaxY-box.MinY+1))
		draw.Draw(region, region.Bounds(), img, image.Point{X: box.MinX, Y: box.MinY}, draw.Src)
		regions[i] = region
	}

	result := make([]int, 0, len(regions))

	for _, region := range regions {
		probabilities := recognizeEnitity(region, models)

		maxIndex := 0

		for modelIndex, value := range probabilities {
			if value > probabilities[maxIndex] {
				maxIndex = modelIndex
			}
		}

		result = append(result, maxIndex)

	}

	return result
}
