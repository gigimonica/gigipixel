package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	drawing := [][]int{
		0: {1, 0, 1, 0},
		1: {0, 1, 1, 1},
	}

	pixelWidth := 20
	pixelHeight := pixelWidth
	pixelColor := color.Black
	width := len(drawing[0]) * pixelWidth
	height := len(drawing) * pixelHeight

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i, row := range drawing {
		for j, pixel := range row {
			if pixel == 1 {
				x := j * pixelWidth
				y := i * pixelHeight
				draw.Draw(img,
					image.Rect(x, y, x+pixelWidth, y+pixelHeight), image.NewUniform(pixelColor), image.ZP, draw.Over)
			}
		}
	}
	f, err := os.Create("gigipixel.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}
