package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
)

func main() {
	// Open a webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// Load the classifier
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("classifier.xml") {
		fmt.Println("Error loading classifier")
		return
	}

	// Create a window to display the webcam
	window := gocv.NewWindow("Webcam")
	defer window.Close()

	// Read frames from the webcam in a loop
	for {
		// Read a frame from the webcam
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Error reading image from webcam")
			return
		}

		// Detect faces in the frame
		rects := classifier.DetectMultiScale(img)

		// Draw a rectangle around each face
		green := color.RGBA{0, 255, 0, 0}
		for _, r := range rects {
			gocv.Rectangle(&img, r, green, 3)
		}

		// Show the webcam frame in the window
		window.IMShow(img)

		// Wait for a key press (30ms)
		if window.WaitKey(30) >= 0 {
			break
		}
	}
}
