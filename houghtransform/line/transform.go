package line

import (
	"image/color"
	"image/draw"
	"log"
	"math"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type DetectedLineParams struct {
	Distance int
	Theta    int
}

type localMaxima2D struct {
	x, y  int
	value uint32
}

func generateEmptyVotes(longestDistance int, degrees int) [][]uint32 {
	result := make([][]uint32, longestDistance, longestDistance)

	for i := 0; i < len(result); i++ {
		result[i] = make([]uint32, degrees, degrees)
	}

	return result

}

func findLocalMaxima(array *[][]uint32, threshold uint32) []localMaxima2D {
	var maximas []localMaxima2D

	for i := 0; i < len(*array); i++ {
		for j := 0; j < len((*array)[i]); j++ {
			value := (*array)[i][j]
			flag := true

			if value <= threshold {
				continue
			}

			if (i-1 >= 0 && j-1 >= 0 && value < (*array)[i-1][j-1]) ||
				(i-1 >= 0 && value < (*array)[i-1][j]) ||
				(j-1 >= 0 && value < (*array)[i][j-1]) ||
				(i+1 < len(*array) && value < (*array)[i+1][j]) ||
				(j+1 < len((*array)[i]) && value < (*array)[i][j+1]) ||
				(i+1 < len(*array) && j+1 < len((*array)[i]) && value < (*array)[i+1][j+1]) ||
				(i-1 >= 0 && j+1 < len((*array)[i]) && value < (*array)[i-1][j+1]) ||
				(j-1 >= 0 && i+1 < len(*array) && value < (*array)[i+1][j-1]) {
				flag = false
			}

			if flag {
				maximas = append(maximas, localMaxima2D{x: i, y: j, value: value})
			}

		}
	}

	return maximas

}

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

func vote(votes *[][]uint32, longestDistance, x, y int, step uint) {
	for theta := uint(0); theta < 180; theta += step {
		thetaRadians := float64(theta) / 180 * math.Pi
		distance := int(math.Round(float64(x)*math.Cos(thetaRadians) + float64(y)*math.Sin(thetaRadians)))
		distance = codeDistance(distance, longestDistance)

		atomic.AddUint32(&(*votes)[distance][int(theta)], 1)
	}

}

func votingRoutine(wg *sync.WaitGroup, img draw.Image, votes *[][]uint32, longestDistance, from, to int) {
	defer wg.Done()
	bounds := img.Bounds()

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA).R

			if pixel == 255 {
				vote(votes, longestDistance, x, y, 1)
			}

		}
	}
}

func HoughTransform(img draw.Image, threshold int, maximumLines int) []DetectedLineParams {
	start := time.Now()

	bounds := img.Bounds()

	// Creating empty matrix quantized by [0deg, 180deg) and [-d; d]
	longestDistance := int(math.Sqrt(float64(bounds.Dx()*bounds.Dx()+bounds.Dy()*bounds.Dy()))) + 1
	votes := generateEmptyVotes(longestDistance*2+1, 180)

	// Voting process
	// Since its calculation heavy, make it concurrent
	threads := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup

	rowsPerThread := bounds.Dy() / threads
	for i := 0; i < threads; i++ {
		wg.Add(1)

		start := img.Bounds().Min.Y + rowsPerThread*i

		var end int

		if i == threads-1 {
			end = img.Bounds().Max.Y

		} else {
			end = start + rowsPerThread

		}

		go votingRoutine(&wg, img, &votes, longestDistance, start, end)
	}
	wg.Wait()

	// Find local maximums, they represent the most precise params for line
	var params []DetectedLineParams
	maxs := findLocalMaxima(&votes, uint32(threshold))

	votes = nil

	// Sort lines by votes count
	sort.Slice(maxs, func(i, j int) bool {
		return maxs[i].value > maxs[j].value
	})

	// First N lines are result
	for i := 0; i < maximumLines && i < len(maxs); i++ {
		value := maxs[i]
		params = append(params, DetectedLineParams{Distance: decodeDistance(value.x, longestDistance), Theta: value.y})
	}

	log.Println("Hough transform:", time.Since(start))

	return params
}
