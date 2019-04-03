package houghtransform

import (
	"image/color"
	"image/draw"
	"log"
	"math"
)

type paramsPair struct {
	distance int
	theta    float64
}

func covertToInterceptForm(distance int, theta float64) func(x int) int {
	theta = 2*math.Pi - theta
	m := -1 * math.Cos(theta) / math.Sin(theta)
	c := float64(distance) * (1 / math.Sin(theta))

	log.Println(m)
	log.Println(c)

	return func(x int) int {
		return int(m*float64(x) + c)
	}

}

func vote(votes *map[paramsPair]uint, x int, y int, step float64) {
	for theta := 0.0; theta <= math.Pi; theta += step {
		distance := float64(x)*math.Cos(theta) - float64(y)*math.Sin(theta)

		// ignores about ±2 deg
		thetaRounded := float64(int(theta*1000)) / 1000
		// down to int here to "group" votes, ± 5 pixels treated as one line
		key := paramsPair{distance: int(distance) / 5 * 5, theta: thetaRounded}

		if val, ok := (*votes)[key]; ok {
			(*votes)[key] = val + 1
		} else {
			(*votes)[key] = 1
		}
	}

}

func HoughTransform(img draw.Image) {
	votes := make(map[paramsPair]uint)

	bounds := img.Bounds()
	hits := 0

	total := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel, _, _, _ := img.At(x, y).RGBA()
			total += 1

			if pixel == 0 {
				hits += 1
				vote(&votes, x, y, math.Pi/180)
			}

		}
	}

	log.Println(hits)
	log.Println(len(votes))

	var result paramsPair
	//

	log.Println("")
	log.Println("--VOTES--")
	for key, val := range votes {
		log.Println(key, val)
	}
	log.Println("--ENDVOTES--")

	for key, val := range votes {

		if val > 800 {
			result = key
		}
	}

	log.Println(result)

	red := color.RGBA{255, 0, 0, 255}
	calc := covertToInterceptForm(result.distance, result.theta)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		y := int(math.Abs(float64(calc(x))))

		if y >= bounds.Min.Y && y < bounds.Max.Y {
			img.Set(x, y, red)
			log.Println(x, y)
		}
	}

}
