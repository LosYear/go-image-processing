package main

import (
	"./baseimage"
	"./rotate"
	"fmt"
	"os"
	"time"
)

/**
Usage:
	<filename> <mode:(lines|circles)> <mode=circles?maxCoin:(1|2|5)>
*/

func main() {
	start := time.Now()

	originalImage, _ := baseimage.ReadFromFile(os.Args[1])

	bounds := originalImage.Bounds()

	originalImage = rotate.Rotate(originalImage, bounds.Min.X, bounds.Max.X, bounds.Min.Y, bounds.Max.Y, 45)

	baseimage.WriteToFile("fixtures/result/output.png", &originalImage, "png")

	fmt.Println("Total took:", time.Since(start))
}
