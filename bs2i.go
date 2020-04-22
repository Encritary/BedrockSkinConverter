package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: bs2i <x size> <y size> <raw skin path>")
		os.Exit(0)
	}

	xSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: bs2i <x size> <y size> <raw skin path>")
		os.Exit(0)
	}

	ySize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usage: bs2i <x size> <y size> <raw skin path>")
		os.Exit(0)
	}

	skinPath := os.Args[3]

	stat, err := os.Stat(skinPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't", err.Error())
		os.Exit(1)
	}

	if stat.Size() != int64(xSize * ySize * 4) {
		fmt.Fprintln(os.Stderr, "Invalid XY size given, size mismatch")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(skinPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't", err.Error())
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(0, 0, xSize, ySize))
	img.Pix = bytes

	basePath := path.Base(skinPath)
	outputPath := basePath[:len(basePath)-len(path.Ext(skinPath))] + ".png"

	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't", err.Error())
		os.Exit(1)
	}

	if err := png.Encode(output, img); err != nil {
		fmt.Fprintln(os.Stderr, "Can't", err.Error())
		os.Exit(1)
	}

	if err := output.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "Can't", err.Error())
		os.Exit(1)
	}
}
