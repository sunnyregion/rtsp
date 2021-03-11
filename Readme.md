# gocv 读取ip Camera 并进行人脸检测

## 安装环境

具体方法参考<https://github.com/hybridgroup/gocv>



这里面也有相关的读取方法，但是这个主要是讲解使用USB摄像头的方式，我们主要讲解连接海康、大华摄像头的方法



## 程序讲解



```go
package main

import (
	"fmt"
	"image/color"
	"gocv.io/x/gocv"
)

func main() {
    // 连接摄像头
	webcam, err := gocv.OpenVideoCapture(`rtsp://admin:password*@10.11.32.198:554/h264/ch33/main/av_stream`)
	if err != nil {
		fmt.Println(`🥵 🥵 🥵 🥵 🥵`, err)
		return
	}
	defer webcam.Close()
	window := gocv.NewWindow("🧒 人脸识别👦")
	defer window.Close()
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected 画人脸框
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	// 打开人脸抠图模型
	if !classifier.Load("./data/haarcascade_frontalface_default.xml") {
		fmt.Println("🥵 🥵 🥵 🥵 🥵Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	for { // 死循环开始取数据
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
            // 出现错误
			fmt.Println(" 👹 👺🥵 🥵 🥵 🥵 🥵读取不到图片🧑🏻‍🚒 👨🏻‍🚒 🥵 🥵 🥵 🥵 🥵👹 👺")
			webcam, _ = gocv.OpenVideoCapture(`rtsp://admin:Cloud1688*@10.11.32.198:554/h264/ch33/main/av_stream`)
			gocv.WaitKey(50)
		}
	}
}

```



Sunny

| 版本 | 时间 |
| ---- | ---- |
|V0.1.0|2021-03-12|