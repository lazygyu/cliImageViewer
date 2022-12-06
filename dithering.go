package main

import (
	"fmt"
	"image"
	"image/color"
)

func Dithering(img image.Image) image.Image {
	var oldPixel, newPixel byte
	var error int
	var idx int
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	origin := toByteArray(img)
	result := image.NewGray(image.Rectangle{Min: image.Pt(0, 0), Max: image.Pt(width, height)})

	var min, max, half byte

	min = 255
	max = 0
	half = 127

	for i := 0; i < width*height; i++ {
		if origin[i] < min {
			min = origin[i]
		}
		if origin[i] > max {
			max = origin[i]
		}
	}

	half = byte((int(max) + int(min)) / 2)

	fmt.Printf("min: %d , max: %d , half: %d\n", min, max, half)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx = (y * width) + x
			oldPixel = origin[y*width+x]
			newPixel = findColor(oldPixel, half)
			error = (int(oldPixel) - int(newPixel)) >> 3
			result.Pix[idx] = newPixel
			if x < width-1 {
				origin[y*width+x+1] = byte(int(origin[y*width+x+1]) + error)
				if x < width-2 {
					origin[y*width+x+2] = byte(int(origin[y*width+x+2]) + error)
				}
			}
			if y < height-1 {
				origin[(y+1)*width+x] = byte(int(origin[(y+1)*width+x]) + error)
			}
		}
	}
	return result
}

func toByteArray(img image.Image) []byte {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	origin := make([]byte, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			origin[y*width+x] = color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y
		}
	}
	return origin
}

func findColor(c, half byte) byte {
	if c > half {
		return 255
	} else {
		return 0
	}
}
