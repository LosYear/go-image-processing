package line

import (
	".."
	"image/draw"
	"log"
	"sort"
	"time"
)

func HoughLineTransform(img draw.Image, threshold int, maximumLines int) []DetectedLineParams {
	start := time.Now()

	// Initialize voting process
	votingProcess := NewLinesVotingProcess(img, 1)
	houghtransform.StartVotingProcess(votingProcess)

	// Find local maximums, they represent the most precise params for line
	var params []DetectedLineParams
	maxs := findLocalMaxima(&votingProcess.votes, uint32(threshold))

	// Sort lines by votes count
	sort.Slice(maxs, func(i, j int) bool {
		return maxs[i].value > maxs[j].value
	})

	// First N lines are result
	for i := 0; i < maximumLines && i < len(maxs); i++ {
		value := maxs[i]
		params = append(params, DetectedLineParams{Distance: decodeDistance(value.x, votingProcess.longestDistance), Theta: value.y})
	}

	log.Println("Hough transform:", time.Since(start))

	return params
}
