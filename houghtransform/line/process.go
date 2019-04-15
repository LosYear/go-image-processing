package line

import (
	"../../helpers"
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync/atomic"
)

type LinesVotingProcess struct {
	img   draw.Image
	votes [][]uint32

	longestDistance int
	step            uint

	sin []float64
	cos []float64
}

func NewLinesVotingProcess(img draw.Image, step uint) LinesVotingProcess {
	process := LinesVotingProcess{img: img, step: step}

	// Initialize empty votes
	bounds := img.Bounds()
	longestDistance := int(math.Sqrt(float64(bounds.Dx()*bounds.Dx()+bounds.Dy()*bounds.Dy()))) + 1

	process.votes = generateEmptyVotes(longestDistance*2+1, 180)
	process.longestDistance = longestDistance

	// Fill trigonometry tables
	helpers.FillSinTable(&process.sin, 180)
	helpers.FillCosTable(&process.cos, 180)

	return process
}

func (p LinesVotingProcess) Vote(x, y int) {
	step := p.step
	longestDistance := p.longestDistance

	for theta := uint(0); theta < 180; theta += step {

		distance := int(math.Round(float64(x)*p.cos[theta] + float64(y)*p.sin[theta]))
		distance = codeDistance(distance, longestDistance)

		atomic.AddUint32(&p.votes[distance][int(theta)], 1)
	}
}

func (p LinesVotingProcess) ImageBounds() image.Rectangle {
	return p.img.Bounds()
}

func (p LinesVotingProcess) ImageAt(x, y int) color.Color {
	return p.img.At(x, y)
}
