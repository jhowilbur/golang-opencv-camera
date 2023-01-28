package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	// Open a webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// Create an image object
	img := gocv.NewMat()
	defer img.Close()

	// Create a window to display the image
	window := gocv.NewWindow("Webcam")
	defer window.Close()

	// Continuously read frames from the webcam and display them in the window
	for {
		// Read a frame from the webcam
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Error reading frame from webcam")
			return
		}

		// Show the frame in the window
		window.IMShow(img)

		// Wait for a key press (30ms)
		if window.WaitKey(30) >= 0 {
			break
		}
	}
}
