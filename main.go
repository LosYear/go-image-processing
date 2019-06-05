package main

import (
	"./baseimage"
	"./imagefilter"
	"./recognition"
	"./recognition/model"
	"fmt"
	"os"
	"strconv"
)

/**
Usage:
	<filename>
*/

func main() {

	// Load models

	models := []model.VectorModel{
		model.Load("./fixtures/models/0.json"),
		model.Load("./fixtures/models/1.json"),
		model.Load("./fixtures/models/2.json"),
		model.Load("./fixtures/models/3.json"),
		model.Load("./fixtures/models/4.json"),
		model.Load("./fixtures/models/5.json"),
		model.Load("./fixtures/models/6.json"),
		model.Load("./fixtures/models/7.json"),
		model.Load("./fixtures/models/8.json"),
		model.Load("./fixtures/models/9.json"),
	}

	for index, m := range models {
		render := model.DrawModel(m, 10, 20)
		baseimage.WriteToFile("fixtures/result/models/"+strconv.Itoa(index)+".png", &render, "png")
	}

	originalImage, _ := baseimage.ReadFromFile(os.Args[1])

	originalImage = imagefilter.AdaptiveThreshold(originalImage, 12, 7)
	baseimage.WriteToFile("fixtures/result/img_thresholded.png", &originalImage, "png")

	fmt.Println(recognition.Recognize(originalImage, models))
}
