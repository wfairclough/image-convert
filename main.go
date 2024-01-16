package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/sunshineplan/imgconv"
)

type OutputSize int

const (
	XXLarge OutputSize = 3840
	XLarge  OutputSize = 1920
	Large   OutputSize = 1280
	Medium  OutputSize = 640
	Small   OutputSize = 320
	XSmall  OutputSize = 200
)

func main() {
	// Read image from file
	src, err := imgconv.Open("testdata/living.jpg")
	if err != nil {
		fmt.Println(err)
	}

	timestamp := time.Now().Format("20060102150405")

	file_writer, _ := os.Create(fmt.Sprintf("output/living-org-%s.pdf", timestamp))
	err = imgconv.Write(file_writer, src, &imgconv.FormatOption{Format: imgconv.PDF})
	if err != nil {
		fmt.Println(err)
	}

	file_writer, _ = os.Create(fmt.Sprintf("output/living-xl-%s.jpg", timestamp))
	err = imgconv.Write(file_writer, resize(src, XLarge), &imgconv.FormatOption{Format: imgconv.JPEG})

	file_writer, _ = os.Create(fmt.Sprintf("output/living-lg-%s.jpg", timestamp))
	err = imgconv.Write(file_writer, resize(src, Large), &imgconv.FormatOption{Format: imgconv.JPEG})

	file_writer, _ = os.Create(fmt.Sprintf("output/living-md-%s.jpg", timestamp))
	err = imgconv.Write(file_writer, resize(src, Medium), &imgconv.FormatOption{Format: imgconv.JPEG})

	file_writer, _ = os.Create(fmt.Sprintf("output/living-sm-%s.jpg", timestamp))
	err = imgconv.Write(file_writer, resize(src, Small), &imgconv.FormatOption{Format: imgconv.JPEG})

	file_writer, _ = os.Create(fmt.Sprintf("output/living-xs-%s.jpg", timestamp))
	err = imgconv.Write(file_writer, resize(src, XSmall), &imgconv.FormatOption{Format: imgconv.JPEG})
}

func resize(src image.Image, size OutputSize) image.Image {
	width := int(size)
	img_size := src.Bounds()
	if img_size.Dx() < int(size) {
		width = img_size.Dx()
	}
	height := width * img_size.Dy() / img_size.Dx()

	return imgconv.Resize(src, &imgconv.ResizeOption{Width: width, Height: height})
}
