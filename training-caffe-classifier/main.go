package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

func main() {
	// Open a webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// Read the Caffe model
	net := gocv.ReadNetFromCaffe("models/google.prototxt", "models/trained.caffemodel")

	if net.Empty() {
		fmt.Println("Error reading network model")
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

		// Convert the frame to a blob
		blob := gocv.BlobFromImage(img, 1.0, image.Pt(224, 224), gocv.NewScalar(104, 117, 123, 0), false, false)

		// Set the blob as the input to the network
		net.SetInput(blob, "data")

		// Run a forward pass to get the classifier's predictions
		prob := net.Forward("prob")

		// Print the predictions
		fmt.Println(prob)

		// Show the webcam frame in the window
		window.IMShow(img)

		// Wait for a key press (30ms)
		if window.WaitKey(30) >= 0 {
			break
		}
	}
}
