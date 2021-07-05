package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	reSize()
}

func genImage() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	f, err := os.Create("demo.jpeg")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(f, m, nil)
	if err != nil {
		panic(err)
	}
}

func readImage() {
	f, err := os.Open("ubuntu.png")
	if err != nil {
		panic(err)
	}
	// decode图片
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	rgba := m.(*image.RGBA)
	subImage := rgba.SubImage(image.Rect(0, 0, 266, 133)).(*image.RGBA)

	// 保存图片
	create, _ := os.Create("new.png")
	err = png.Encode(create, subImage)
	if err != nil {
		panic(err)
	}
}

func b64() {
	f, err := os.Open("ubuntu.png")
	if err != nil {
		panic(err)
	}
	all, _ := ioutil.ReadAll(f)

	str := base64.StdEncoding.EncodeToString(all)

	// 注意，上述的结果并不包含头部分，所以得自己加一个 data:image/png;base64,
	fmt.Printf("%s\n", str)
}

func reSize() {
	// open "test.jpg"
	file, err := os.Open("ubuntu.png")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(100, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
