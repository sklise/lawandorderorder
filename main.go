package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/otiai10/gosseract"
)

const (
	thresh = 0.75
)

func blackPixelPercent(im image.Image) float64 {
	bounds := im.Bounds()

	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	pixelCount := (bounds.Max.Y - bounds.Min.Y) * (bounds.Max.X - bounds.Min.X)

	blackPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := im.At(x, y).RGBA()

			if r == 0 && g == 0 && b == 0 {
				blackPixels++
			}
		}
	}
	return float64(blackPixels) / float64(pixelCount)
}

func loadImage(s string) (image.Image, error) {
	reader, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	return m, err
}

func main() {
	// Decode the JPEG data. If reading from file, create a reader with

	season := "data/frames/16/"
	// names of folders of episodes
	episodeDirs, _ := ioutil.ReadDir(season)
	for e := range episodeDirs {
		if episodeDirs[e].IsDir() {
			// names of files
			episodeName := episodeDirs[e].Name()
			episode, _ := ioutil.ReadDir(season + episodeName)

			for i := range episode {
				m, err := loadImage(season + episodeName + "/" + episode[i].Name())
				if err != nil {
					log.Fatal("Error loading image.")
				}

				blackPixels := blackPixelPercent(m)
				if blackPixels > thresh {
					out := gosseract.Must(gosseract.Params{Src: season + episodeName + "/" + episode[i].Name(), Languages: "eng"})

					if out != "" {
						fmt.Println("candidate frame", episode[i].Name())
						fmt.Println(out)
					}
				}
			}
		}
	}
}
