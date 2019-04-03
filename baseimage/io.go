package baseimage

import (
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func ReadFromFile(filename string) (draw.Image, string) {
	reader, err := os.Open(filename)

	defer func() {
		err := reader.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(reader)
	drawImage := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(drawImage, drawImage.Bounds(), img, img.Bounds().Min, draw.Src)

	if err != nil {
		log.Fatal(err)
	}

	return drawImage, format
}

func WriteToFile(filename string, img *draw.Image, format string) {
	writer, err := os.Create(filename)
	defer func() {
		err := writer.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	if format == "jpeg" {
		err := jpeg.Encode(writer, *img, nil)

		if err != nil {
			log.Fatal(err)
		}
	} else if format == "png" {
		err := png.Encode(writer, *img)

		if err != nil {
			log.Fatal(err)
		}
	}
}
