package rotate

import (
	"math"
)

func mapDestinationToSource(x, y, centerX, centerY, angle int) (int, int) {
	fAngle := float64(angle) * math.Pi / 180

	mappedX := math.Round(float64(centerX) + float64(x-centerX)*math.Cos(fAngle) + float64(y-centerY)*math.Sin(fAngle))
	mappedY := math.Round(float64(centerY) - float64(x-centerX)*math.Sin(fAngle) + float64(y-centerY)*math.Cos(fAngle))

	return int(mappedX), int(mappedY)
}
