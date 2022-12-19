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
	// Open an image file
	f, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Decode the image
	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Create a new image with the same size as the input image
	outputImg := image.NewRGBA(img.Bounds())

	// Draw a red rectangle on the output image
	draw.Draw(outputImg, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)

	// Open an output file
	f, err = os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// Encode the output image and write it to the output file
	if err := png.Encode(f, outputImg); err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}
}
