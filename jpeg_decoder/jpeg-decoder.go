package jpeg_decoder

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io"
	"log"
	"os"
	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
)

type RGBASum struct {
	r, g, b, a uint64
}

func AvgColorFromFile(filename string) (avg color.Color, err error) {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	return AvgColor(reader)
}

func AvgColor(reader io.Reader) (avg color.Color, err error) {
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var total uint64
	total = 0
	sum := RGBASum{}

	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			total++
			r, g, b, a := m.At(x, y).RGBA()
			sum.r = (sum.r + uint64(r))
			sum.g = (sum.g + uint64(g))
			sum.b = (sum.b + uint64(b))
			sum.a = (sum.a + uint64(a))
		}
	}
	r := sum.r / total
	g := sum.g / total
	b := sum.b / total
	a := sum.a / total
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	// reader, err := os.Open("testdata/video-001.q50.420.jpeg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	reader, err := os.Open("barracks.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}
	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
