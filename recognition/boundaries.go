package recognition

import (
	"image"
	"image/color"
	"image/draw"
)

type BoundingBox struct {
	MinX, MaxX, MinY, MaxY int
}

type minMaxPair struct {
	min, max int
	hasData  bool
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindBoundingBoxes(img draw.Image) []BoundingBox {
	imgBoundaries := img.Bounds()
	projection := make([]minMaxPair, imgBoundaries.Max.X-imgBoundaries.Min.X+1, imgBoundaries.Max.X-imgBoundaries.Min.X+1)

	for x := imgBoundaries.Min.X; x < imgBoundaries.Max.X; x++ {
		for y := imgBoundaries.Min.Y; y < imgBoundaries.Max.Y; y++ {
			pixel := img.At(x, y).(color.RGBA).R

			if pixel != 255 {
				if !projection[x].hasData {
					projection[x].min = imgBoundaries.Max.X + 1
				}

				projection[x].hasData = true
				projection[x].max = max(projection[x].max, y)
				projection[x].min = min(projection[x].min, y)
			}
		}
	}

	projection = append([]minMaxPair{{0, 0, false}}, projection...)

	tmp := BoundingBox{0, 0, 0, 0}

	boundaries := make([]BoundingBox, 0)

	for i := 1; i < len(projection); i++ {
		if !projection[i-1].hasData && projection[i].hasData {
			tmp.MinX = i - 1
			tmp.MinY = projection[i].min
			tmp.MaxY = projection[i].max
			tmp.MaxX = i
		} else if projection[i-1].hasData && projection[i].hasData {
			tmp.MaxX = i - 1
			tmp.MinY = min(tmp.MinY, projection[i].min)
			tmp.MaxY = max(tmp.MaxY, projection[i].max)
		} else if projection[i-1].hasData && !projection[i].hasData {
			boundaries = append(boundaries, tmp)
		}
	}

	return boundaries
}

func DrawBoundaries(img draw.Image, boundaries []BoundingBox) draw.Image {
	imgBounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, imgBounds.Dx(), imgBounds.Bounds().Dy()))
	draw.Draw(result, imgBounds, img, imgBounds.Min, draw.Src)

	for _, boundary := range boundaries {
		red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

		result.Set(boundary.MinX, boundary.MinY, red)
		result.Set(boundary.MinX, boundary.MaxY, red)
		result.Set(boundary.MaxX, boundary.MinY, red)
		result.Set(boundary.MaxX, boundary.MaxY, red)
	}

	return result
}
