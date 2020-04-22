package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: i2bs <image1> [image2] ... [imageN]")
		os.Exit(0)
	}

	imagePaths := os.Args[1:]
	for _, imagePath := range imagePaths {
		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't", err.Error())
			os.Exit(1)
		}

		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't open", imagePath, err.Error())
			os.Exit(1)
		}

		basePath := path.Base(imagePath)
		outputPath := basePath[:len(basePath)-len(path.Ext(imagePath))] + ".skindata"

		output, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't", err.Error())
			os.Exit(1)
		}

		bounds := img.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				_, err := output.Write([]byte{uint8(r), uint8(g), uint8(b), uint8(a)})

				if err != nil {
					fmt.Fprintln(os.Stderr, "Can't", err.Error())
					os.Exit(1)
				}
			}
		}

		if err := output.Close(); err != nil {
			fmt.Fprintln(os.Stderr, "Can't", err.Error())
			os.Exit(1)
		}
	}
}
