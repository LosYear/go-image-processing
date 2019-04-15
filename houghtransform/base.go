package houghtransform

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

type VotingProcess interface {
	Vote(x, y int)

	ImageBounds() image.Rectangle
	ImageAt(x, y int) color.Color
}

func votingRoutine(wg *sync.WaitGroup, process *VotingProcess, from, to int) {
	defer wg.Done()
	bounds := (*process).ImageBounds()

	for y := from; y < to; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := (*process).ImageAt(x, y).(color.RGBA).R

			if pixel == 255 {
				(*process).Vote(x, y)
			}

		}
	}
}

func StartVotingProcess(process VotingProcess) {
	bounds := process.ImageBounds()

	// Voting process
	// Since its calculation is heavy, make it concurrent
	threads := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup

	rowsPerThread := bounds.Dy() / threads
	for i := 0; i < threads; i++ {
		wg.Add(1)

		start := bounds.Min.Y + rowsPerThread*i

		var end int

		if i == threads-1 {
			end = bounds.Max.Y

		} else {
			end = start + rowsPerThread

		}

		go votingRoutine(&wg, &process, start, end)
	}
	wg.Wait()
}
