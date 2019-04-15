package coins

import (
	"log"
	"math"
)
import "../houghtransform/circle"

const fiveRubles = 145.0
const twoRubles = 132.0
const oneRuble = 120.0

func CountMoney(coins []circle.DetectedCircleParams, maxCoin int, precision float64) [3]int {
	var maxCanonicalRadius float64
	if maxCoin == 1 {
		maxCanonicalRadius = oneRuble
	} else if maxCoin == 2 {
		maxCanonicalRadius = twoRubles

	} else if maxCoin == 5 {
		maxCanonicalRadius = fiveRubles

	} else {
		log.Fatal("Unsupported coin type")
	}

	ratios := [3]float64{oneRuble / maxCanonicalRadius, twoRubles / maxCanonicalRadius, fiveRubles / maxCanonicalRadius}

	result := [3]int{0, 0, 0}

	// Find maximum radius
	maxRadius := 0
	for _, c := range coins {
		if maxRadius < c.Radius {
			maxRadius = c.Radius
		}
	}

	// Recognize coins
	for _, c := range coins {
		ratio := float64(c.Radius) / float64(maxRadius)

		for dignity, idealRatio := range ratios {
			if math.Abs(ratio-idealRatio) <= precision {
				result[dignity]++
				break
			}
		}
	}

	return result
}
