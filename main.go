package main

import (
	"./baseimage"
	"./coins"
	houghtransformcircle "./houghtransform/circle"
	houghtransformLine "./houghtransform/line"
	houghVisualization "./houghtransform/visualization"
	"./imagefilter"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/**
Usage:
	<filename> <mode:(lines|circles)> <mode=circles?maxCoin:(1|2|5)>
*/

func main() {
	start := time.Now()

	mode := os.Args[2]
	originalImage, _ := baseimage.ReadFromFile(os.Args[1])
	img := imagefilter.CannyEdgeDetection(originalImage)
	if mode == "lines" {
		linesSet := houghtransformLine.HoughLineTransform(img, 150, 1)
		log.Println("Objects detected:", len(linesSet))
		originalImage = houghVisualization.DrawHoughLinesSet(originalImage, linesSet)
	} else if mode == "circles" {
		maxCoin, _ := strconv.ParseInt(os.Args[3], 10, 64)
		circleSet := houghtransformcircle.HoughCircleTransform(img, 150, 50)
		circleSet = coins.RemoveRedundantCircles(circleSet, 10)
		originalImage = houghVisualization.DrawHoughCirclesSet(originalImage, circleSet)

		coinsCount := coins.CountMoney(circleSet, int(maxCoin), 0.04)
		sum := 0

		for dignity, count := range coinsCount {
			if dignity == 0 {
				sum += count
				fmt.Println("One ruble:", count)
			} else if dignity == 1 {
				sum += count * 2
				fmt.Println("Two rubles:", count)
			} else {
				sum += count * 5
				fmt.Println("Five rubles:", count)
			}
		}
		fmt.Println("Sum:", sum)
		fmt.Println("")
	} else {
		log.Fatal("Unsupported mode")
	}

	baseimage.WriteToFile("fixtures/result/skeleton.png", &img, "png")
	baseimage.WriteToFile("fixtures/result/output.png", &originalImage, "png")

	fmt.Println("Total took:", time.Since(start))
}
