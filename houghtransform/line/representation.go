package line

import "math"

func ConvertToInterceptForm(distance int, theta float64) func(x int) float64 {
	theta = 2*math.Pi - theta
	m := -1 * math.Cos(theta) / math.Sin(theta)
	c := float64(distance) * (1 / math.Sin(theta))

	return func(x int) float64 {
		return m*float64(x) + c
	}

}

func codeDistance(distance, longestDistance int) int {
	return distance + longestDistance
}

func decodeDistance(distance, longestDistance int) int {
	return distance - longestDistance
}
