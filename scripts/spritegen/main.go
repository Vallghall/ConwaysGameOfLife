package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const size = 10

func main() {
	f, err := os.Open("../../cmd/configs/configs.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	aliveImage := image.NewRGBA(image.Rect(0, 0, size, size))
	deadImage := image.NewRGBA(image.Rect(0, 0, size, size))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			aliveImage.Set(i, j, color.Black)
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j > 0 && j < size-1 && i > 0 && i < size-1 {
				deadImage.Set(i, j, color.White)
				continue
			}
			deadImage.Set(i, j, color.Gray16{Y: 0xaaaa})
		}
	}

	f1, err := os.Create("../../cmd/configs/alive.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err = png.Encode(f1, aliveImage); err != nil {
		log.Fatalln(err)
	}

	f2, err := os.Create("../../cmd/configs/dead.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err = png.Encode(f2, deadImage); err != nil {
		log.Fatalln(err)
	}
}
