package main

import (
	"./baseimage"
	"./houghtransform"
	houghtransformLine "./houghtransform/line"
	"./imagefilter"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	originalImage, _ := baseimage.ReadFromFile(os.Args[1])
	img := imagefilter.CannyEdgeDetection(originalImage)
	linesSet := houghtransformLine.HoughTransform(img, 150, 50)
	log.Println("Objects detected:", len(linesSet))
	originalImage = houghtransform.DrawHoughLinesSet(originalImage, linesSet)
	baseimage.WriteToFile("fixtures/test.png", &originalImage, "png")

	fmt.Println("Total took:", time.Since(start))
}
