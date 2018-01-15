package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// this is the conversion of our drawing
	drawing := [][]int{
		0: {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		1: {0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0},
		2: {0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		3: {0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		4: {0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		5: {0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		6: {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		7: {0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		8: {0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		9: {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Setup the drawing style
	pixelWidth := 100
	pixelHeight := 100

	pixelColor := color.Black
	width := len(drawing[0]) * pixelWidth
	height := len(drawing) * pixelHeight

	// define the image we are going to draw on
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// white background
	draw.Draw(img,
		image.Rect(0, 0, width, height),
		image.NewUniform(color.White),
		image.ZP, draw.Over)

	// convert our 1 and 0 into pixels
	for i, row := range drawing {
		for j, pixel := range row {
			if pixel == 1 {
				// calculate the coordinates where to start the drawing
				x := j * pixelWidth
				y := i * pixelHeight
				// calculate where to end the "pixel"
				xEnd := x + pixelWidth
				yEnd := y + pixelHeight
				// draw the "pixel" on top of the image
				draw.Draw(img,
					image.Rect(x, y, xEnd, yEnd),
					image.NewUniform(pixelColor),
					image.ZP, draw.Over)
			}
		}
	}
	// create a new file to save our image
	f, err := os.Create("gigipixel.png")
	if err != nil {
		panic(err)
	}
	// save our image as a png
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	// close the file we opened
	f.Close()
	fmt.Println("open gigipixel.png")
}
