package imageconvert

import (
	"fmt"
	"image"
	"os"
	"time"
  "path/filepath"

	"github.com/sunshineplan/imgconv"
)

type OutputSize int

func (s OutputSize) String() string {
  switch s {
  case XXLarge:
    return "xxl"
  case XLarge:
    return "xl"
  case Large:
    return "lg"
  case Medium:
    return "md"
  case Small:
    return "sm"
  case XSmall:
    return "xs"
  default:
    return "unknown"
  }
}

func (s OutputSize) Width() int {
  return int(s)
}

func (s OutputSize) Values() []OutputSize {
  return []OutputSize{XXLarge, XLarge, Large, Medium, Small, XSmall}
}

const (
	XXLarge OutputSize = 3840
	XLarge  OutputSize = 1920
	Large   OutputSize = 1280
	Medium  OutputSize = 640
	Small   OutputSize = 320
	XSmall  OutputSize = 200
)

// func main() {
// 	src, err := imgconv.Open("testdata/living.jpg")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
//   err = write_image(src, XXLarge, imgconv.PDF)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
//   for _, size := range XXLarge.Values() {
//     err = write_image(src, size, imgconv.JPEG)
//     if err != nil {
//       fmt.Println(err)
//     }
//   }
// }

func WriteResizedImage(src image.Image, size OutputSize, format imgconv.Format) error {
  timestamp := time.Now().Format("200601021504")
  output_path := filepath.Join("output", fmt.Sprintf("living-%s-%s.%s", size.String(), timestamp, format.String()))
  file_writer, _ := os.Create(output_path)
  return imgconv.Write(file_writer, resize(src, size), &imgconv.FormatOption{Format: format})
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

