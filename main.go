package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.OpenVideoCapture(`rtsp://admin:password*@10.11.32.198:554/h264/ch33/main/av_stream`)
	if err != nil {
		fmt.Println(`π₯΅ π₯΅ π₯΅ π₯΅ π₯΅`, err)
		return
	}
	defer webcam.Close()
	window := gocv.NewWindow("π§ δΊΊθΈθ―ε«π¦")
	defer window.Close()
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("./data/haarcascade_frontalface_default.xml") {
		fmt.Println("π₯΅ π₯΅ π₯΅ π₯΅ π₯΅Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	for {
		if webcam.Read(&img) {
			if img.Empty() {
				continue
			}
			// detect faces
			rects := classifier.DetectMultiScale(img)
			//fmt.Printf("found %d faces\n", len(rects))

			// draw a rectangle around each face on the original image
			for _, r := range rects {
				gocv.Rectangle(&img, r, blue, 3)
			}
			window.IMShow(img)
			window.WaitKey(1)
		} else {
			fmt.Println(" πΉ πΊπ₯΅ π₯΅ π₯΅ π₯΅ π₯΅θ―»εδΈε°εΎηπ§π»βπ π¨π»βπ π₯΅ π₯΅ π₯΅ π₯΅ π₯΅πΉ πΊ")
			webcam, _ = gocv.OpenVideoCapture(`rtsp://admin:Cloud1688*@10.11.32.198:554/h264/ch33/main/av_stream`)
			gocv.WaitKey(50)
		}
	}
}
