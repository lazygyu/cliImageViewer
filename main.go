package main

import (
	"fmt"
	"gv/dotscreen"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Usage: gv filename")
	}

	filename := args[0]

	// open file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("cannot open the file")
	}

	defer f.Close()

	// load img
	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal("Cannot decode the image")
	}

	d := dotscreen.New()
	fmt.Printf("screen width: %d, height: %d\n", d.Width, d.Height)

	fmt.Printf("image width: %d, height: %d\n", img.Bounds().Max.X, img.Bounds().Max.Y)

	// resize to fit
	if img.Bounds().Max.X > d.Width && img.Bounds().Max.Y > d.Height {
		xratio := float32(img.Bounds().Max.X) / float32(d.Width)
		yratio := float32(img.Bounds().Max.Y) / float32(d.Height)
		var w, h int
		if xratio >= yratio {
			w = d.Width
			h = int(float32(img.Bounds().Max.Y) / xratio)
		} else {
			w = int(float32(img.Bounds().Max.X) / yratio)
			h = d.Height
		}
		img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
	} else if img.Bounds().Max.X > d.Width {
		ratio := float32(img.Bounds().Max.X) / float32(d.Width)
		w := d.Width
		h := int(float32(img.Bounds().Max.Y) / ratio)
		img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
	} else if img.Bounds().Max.Y > d.Height {
		ratio := float32(img.Bounds().Max.Y) / float32(d.Height)
		w := int(float32(img.Bounds().Max.X) / ratio)
		h := d.Height
		img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
	}
	fmt.Printf("after resize width: %d, height: %d\n", img.Bounds().Max.X, img.Bounds().Max.Y)

	// dithering
	result := Dithering(img)
	// turn into a byte array
	pixels := toByteArray(result)
	// print out
	dotscreen.PrintImage(img.Bounds().Max.X, img.Bounds().Max.Y, &pixels)
}
