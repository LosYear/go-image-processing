package main

import (
	"./baseimage"
	"./imagefilter"
	"fmt"
	"os"
)

func main() {
	img, _ := baseimage.ReadFromFile(os.Args[1])
	img = imagefilter.CannyEdgeDetection(img)
	//houghtransform.HoughTransform(img)
	baseimage.WriteToFile("fixtures/test.png", &img, "png")

	fmt.Println("FINISHED")
}
