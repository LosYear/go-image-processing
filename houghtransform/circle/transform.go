package circle

import (
	".."
	"image/draw"
	"log"
	"sort"
	"time"
)

func HoughCircleTransform(img draw.Image, threshold int, maximumLines int) []DetectedCircleParams {
	start := time.Now()

	// Initialize voting process
	votingProcess := NewCirclesVotingProcess(img, 20, 1)
	houghtransform.StartVotingProcess(votingProcess)

	// Find local maximas
	var params []DetectedCircleParams
	maxs := findLocalMaxima(&votingProcess.votes, uint32(threshold))

	// Sort lines by votes count
	sort.Slice(maxs, func(i, j int) bool {
		return maxs[i].value > maxs[j].value
	})

	// First N lines are result
	for i := 0; i < maximumLines && i < len(maxs); i++ {
		value := maxs[i]
		a, b := decodeCenter(value.x, value.y, votingProcess.maxX, votingProcess.maxY)
		params = append(params, DetectedCircleParams{A: a, B: b, Radius: value.z})
	}

	log.Println("Hough transform:", time.Since(start))

	return params
}
