package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	// Read an image file
	img := gocv.IMRead("./img/image.png", gocv.IMReadColor)

	// Check for error
	if img.Empty() {
		fmt.Println("Error reading image")
		return
	}

	// Create a window to display the image
	window := gocv.NewWindow("Image")
	defer window.Close()

	// Show the image in the window
	window.IMShow(img)

	// Wait for a key press
	window.WaitKey(0)
}
