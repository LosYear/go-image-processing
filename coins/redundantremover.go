package coins

import (
	houghCircleTransform "../houghtransform/circle"
	"math"
	"sort"
)

func RemoveRedundantCircles(circles []houghCircleTransform.DetectedCircleParams, centerDelta uint) []houghCircleTransform.DetectedCircleParams {
	var result []houghCircleTransform.DetectedCircleParams

	sort.Slice(circles, func(i, j int) bool {
		return circles[i].Radius < circles[j].Radius
	})

	for i := 0; i < len(circles); i++ {
		flag := true
		for j := i + 1; j < len(circles); j++ {
			distance := math.Pow(float64(circles[i].A-circles[j].A), 2) + math.Pow(float64(circles[i].B-circles[j].B), 2)

			if math.Sqrt(distance) < float64(centerDelta) {
				flag = false
				break
			}

			// Let's assume that there are no intersecting circles
			comparedRadius := math.Max(float64(circles[j].Radius), float64(circles[i].Radius))
			if distance < comparedRadius*comparedRadius {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, circles[i])
		}
	}

	return result
}
