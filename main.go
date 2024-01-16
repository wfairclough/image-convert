package main

import (
  "fmt"
  "os"
  "time"

  "github.com/sunshineplan/imgconv"
)

func main() {
  // Read image from file
  src, err := imgconv.Open("testdata/living.jpg")
  if err != nil {
    fmt.Println(err)
  }

  file_writer, err := os.Create(fmt.Sprintf("testdata/living-%s.pdf", time.Now().Format("20060102150405")))

  // Convert image from io.Reader to io.Writer
  err = imgconv.Write(file_writer, src, &imgconv.FormatOption{Format: imgconv.PDF})
  if err != nil {
    fmt.Println(err)
  }
}

