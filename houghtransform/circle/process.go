package circle

import (
	"../../helpers"
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync/atomic"
)

type CirclesVotingProcess struct {
	img   draw.Image
	votes [][][]uint32

	minRadius, maxRadius, maxX, maxY int
	step                             uint

	sin []float64
	cos []float64
}

func NewCirclesVotingProcess(img draw.Image, minRadius int, step uint) CirclesVotingProcess {
	process := CirclesVotingProcess{img: img, minRadius: minRadius, step: step}

	// Initialize empty votes
	bounds := img.Bounds()
	maxRadius := int(math.Min(float64(bounds.Dx()), float64(bounds.Dy()))/2 + 1)

	process.votes = generateEmptyVotes(maxRadius, bounds.Max.X*2+1, bounds.Max.Y*2+1)
	process.maxRadius = maxRadius

	// Fill trigonometry tables
	helpers.FillSinTable(&process.sin, 360)
	helpers.FillCosTable(&process.cos, 360)

	return process
}

func (p CirclesVotingProcess) Vote(x, y int) {
	for radius := p.minRadius; radius < p.maxRadius; radius++ {
		for theta := uint(0); theta < 360; theta += p.step {
			a := int(float64(x) - float64(radius)*p.cos[theta])
			b := int(float64(y) - float64(radius)*p.sin[theta])
			a, b = codeCenter(a, b, p.maxX, p.maxY)

			if a >= 0 && b >= 0 && a < len(p.votes) && b < len(p.votes[a]) {
				atomic.AddUint32(&p.votes[a][b][radius], 1)
			}
		}
	}
}

func (p CirclesVotingProcess) ImageBounds() image.Rectangle {
	return p.img.Bounds()
}

func (p CirclesVotingProcess) ImageAt(x, y int) color.Color {
	return p.img.At(x, y)
}
