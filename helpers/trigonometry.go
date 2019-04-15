package helpers

import "math"

func FillSinTable(sins *[]float64, size int) {
	*sins = make([]float64, size, size)
	for theta := 0; theta < size; theta++ {
		thetaRadians := float64(theta) / 180 * math.Pi
		(*sins)[theta] = math.Sin(thetaRadians)
	}
}

func FillCosTable(cos *[]float64, size int) {
	*cos = make([]float64, size, size)

	for theta := 0; theta < size; theta++ {
		thetaRadians := float64(theta) / 180 * math.Pi
		(*cos)[theta] = math.Cos(thetaRadians)
	}
}
