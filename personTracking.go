package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func main() {

	// define default hog descriptor
	hog := gocv.NewHOGDescriptor()
	defer hog.Close()
	hog.SetSVMDetector(gocv.HOGDefaultPeopleDetector())

	// color for the rect when faces detected
	//blue := color.RGBA{0, 0, 255, 0}

	// open webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("error opening video capture device")
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Person Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed")
			return
		}
		if img.Empty() {
			continue
		}

		//resize image
		fact := float64(400) / float64(img.Cols())
		newY := float64(img.Rows()) * fact
		gocv.Resize(img, &img, image.Point{X: 400, Y: int(newY)}, 0, 0, 1)

		// detect person
		rects := hog.DetectMultiScale(img)
		fmt.Printf("found %d persons\n", len(rects))

		// for _, r := range rects {
		// 	gocv.Rectangle(&img, r, blue, 3)

		// 	size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		// 	pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		// 	gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		// }

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}
