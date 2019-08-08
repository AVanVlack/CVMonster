package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

type tracker struct {
	people   []int
	detector *gocv.HOGDescriptor
	input    gocv.VideoCapture
	//viewer   *gocv.Window
	img gocv.Mat
}

func newTracker() *tracker {
	// define default hog descriptor
	PersonDetector := gocv.NewHOGDescriptor()
	defer PersonDetector.Close()
	err := PersonDetector.SetSVMDetector(gocv.HOGDefaultPeopleDetector())
	if err != nil {
		fmt.Printf("error setting SVM Detector")
	}

	// // open webcam
	//webcam, err := gocv.OpenVideoCapture(0)
	// if err != nil {
	// 	fmt.Printf("error opening video capture device")
	// }
	// defer webcam.Close()

	// use file
	file := os.Args[1]
	video, err := gocv.VideoCaptureFile(file)
	if err != nil {
		fmt.Printf("error opening video file")
	}
	defer video.Close()

	// open display window
	//window := gocv.NewWindow("Person Detect")
	//defer window.Close()

	// prepare image matrix
	mat := gocv.NewMat()
	defer mat.Close()

	return &tracker{
		detector: &PersonDetector,
		input:    video,
		//viewer:   window,
		img: mat,
	}
}

func (t *tracker) processFrame() {
	// color for people box
	color := color.RGBA{0, 0, 255, 0}

	// grab new frame from input
	if ok := t.input.Read(&t.img); !ok {
		fmt.Printf("Device closed")
		return
	}
	if t.img.Empty() {
		fmt.Print("hello")
		return
	}
	fmt.Printf("Device closed")

	// resize image
	fact := float64(400) / float64(t.img.Cols())
	newY := float64(t.img.Rows()) * fact
	gocv.Resize(t.img, &t.img, image.Point{X: 400, Y: int(newY)}, 0, 0, 1)

	// detect person
	rects := t.detector.DetectMultiScaleWithParams(t.img, 0, image.Point{X: 8, Y: 8}, image.Point{X: 16, Y: 16}, 1.05, 2, false)
	fmt.Printf("found %d persons\n", len(rects))

	for _, r := range rects {
		gocv.Rectangle(&t.img, r, color, 3)

		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&t.img, "Human", pt, gocv.FontHersheyPlain, 1.2, color, 2)
	}

	// show the image in the window and set peoples location
	//t.viewer.IMShow(t.img)
}

// func trackPersons() {
// 	// color for the rect when faces detected

// 	// define default hog descriptor
// 	hog := gocv.NewHOGDescriptor()
// 	defer hog.Close()
// 	err := hog.SetSVMDetector(gocv.HOGDefaultPeopleDetector())
// 	if err != nil {
// 		fmt.Printf("error setting SVM Detector")
// 		return
// 	}

// 	// color for the rect when faces detected
// 	blue := color.RGBA{0, 0, 255, 0}

// 	// // open webcam
// 	//webcam, err := gocv.OpenVideoCapture(0)
// 	// if err != nil {
// 	// 	fmt.Printf("error opening video capture device")
// 	// 	return
// 	// }
// 	// defer webcam.Close()

// 	//Use file
// 	file := os.Args[1]
// 	webcam, err := gocv.VideoCaptureFile(file)
// 	if err != nil {
// 		fmt.Printf("error opening video file")
// 		return
// 	}
// 	defer webcam.Close()

// 	//open display window
// 	window := gocv.NewWindow("Person Detect")
// 	defer window.Close()

// 	// prepare image matrix
// 	img := gocv.NewMat()
// 	defer img.Close()
// 	for {
// 		if ok := webcam.Read(&img); !ok {
// 			fmt.Printf("Device closed")
// 			return
// 		}
// 		if img.Empty() {
// 			continue
// 		}

// 		//resize image
// 		fact := float64(400) / float64(img.Cols())
// 		newY := float64(img.Rows()) * fact
// 		gocv.Resize(img, &img, image.Point{X: 400, Y: int(newY)}, 0, 0, 1)

// 		// detect person
// 		//rects := hog.DetectMultiScale(img)
// 		rects := hog.DetectMultiScaleWithParams(img, 0, image.Point{X: 8, Y: 8}, image.Point{X: 16, Y: 16}, 1.05, 2, false)
// 		fmt.Printf("found %d persons\n", len(rects))

// 		for _, r := range rects {
// 			gocv.Rectangle(&img, r, blue, 3)

// 			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
// 			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
// 			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
// 		}

// 		// show the image in the window, and wait 1 millisecond
// 		window.IMShow(img)
// 		if window.WaitKey(1) >= 0 {
// 			break
// 		}
// 	}

// }
