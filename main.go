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

func countBlackPixels(im image.Image) int {
	bounds := im.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.

	blackPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := im.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].

			if r == 0 && g == 0 && b == 0 {
				blackPixels++
			}
		}
	}
	return blackPixels
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
	//

	redball, _ := ioutil.ReadDir("data/redball")
	for i := range redball {
		m, err := loadImage("data/redball/" + redball[i].Name())
		if err != nil {
			log.Fatal("Error loading image.")
		}
		bounds := m.Bounds()
		pixelCount := (bounds.Max.Y - bounds.Min.Y) * (bounds.Max.X - bounds.Min.X)
		blackPixels := countBlackPixels(m)
		if float64(blackPixels)/float64(pixelCount) > thresh {
			fmt.Println("candidate frame", redball[i].Name())
			out := gosseract.Must(gosseract.Params{Src: "data/redball/" + redball[i].Name(), Languages: "eng"})
			fmt.Println(out)
		}
	}

}
