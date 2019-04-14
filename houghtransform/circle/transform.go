package circle

import (
	"image/color"
	"image/draw"
	"math"
	"sort"
)

func codeCenter(a, b, maxX, maxY int) (int, int) {
	return a + maxX, b + maxY
}

func decodeCenter(a, b, maxX, maxY int) (int, int) {
	return a - maxX, b - maxY
}

func vote(votes *[][][]uint32, minRadius, maxRadius, maxX, maxY, x, y int, step uint) {
	for radius := minRadius; radius < maxRadius; radius++ {
		for theta := uint(0); theta < 360; theta += step {
			thetaRadians := float64(theta) / 180 * math.Pi
			a := int(float64(x) - float64(radius)*math.Cos(thetaRadians))
			b := int(float64(y) - float64(radius)*math.Sin(thetaRadians))
			a, b = codeCenter(a, b, maxX, maxY)

			if a < len(*votes) && b < len((*votes)[0]) {
				(*votes)[a][b][radius] += 1
			}
		}
	}

}

func HoughCircleTransform(img draw.Image, threshold int, maximumLines int) []DetectedCircleParams {
	//start := time.Now()

	bounds := img.Bounds()

	// Creating empty matrix quantized by [0, min{width, height} / 2 + 1) and [0, dimensions]
	maxRadius := int(math.Min(float64(bounds.Dx()), float64(bounds.Dy()))/2 + 1)
	votes := generateEmptyVotes(maxRadius, bounds.Max.X*2+1, bounds.Max.Y*2+1)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA).R

			if pixel == 255 {
				vote(&votes, 20, maxRadius, bounds.Max.X, bounds.Max.Y, x, y, 1)
			}

		}
	}

	var params []DetectedCircleParams
	maxs := findLocalMaxima(&votes, uint32(threshold))

	votes = nil

	// Sort lines by votes count
	sort.Slice(maxs, func(i, j int) bool {
		return maxs[i].value > maxs[j].value
	})

	// First N lines are result
	for i := 0; i < maximumLines && i < len(maxs); i++ {
		value := maxs[i]
		a, b := decodeCenter(value.x, value.y, bounds.Max.X, bounds.Max.Y)
		params = append(params, DetectedCircleParams{A: a, B: b, Radius: value.z})
	}

	return params
}
