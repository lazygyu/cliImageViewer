package main

import (
	"fmt"
)

var dotPoisitions []uint = []uint{1, 8, 2, 16, 4, 32, 64, 128}

func getPixel(width int, height int, data *[]byte, x int, y int) byte {
	if x >= width || x < 0 {
		return 0
	}

	if y >= height || y < 0 {
		return 0
	}

	return (*data)[y*width+x]
}

func PrintImage(width, height int, data *[]byte, invert bool) {
	var pixel byte
	var ch uint
	var sx, sy, xlimit, ylimit int

	xChunk := width / 2
	yChunk := height / 5
	if width%2 != 0 {
		xChunk++
	}
	if height%5 != 0 {
		yChunk++
	}

	targetPixelValue := byte(0)
	if invert {
		targetPixelValue = 1
	}

	for y := 0; y < yChunk; y++ {
		for x := 0; x < xChunk; x++ {
			ch = 0
			sx = x * 2
			sy = y * 5
			xlimit = min(2, width-sx)
			ylimit = min(4, height-sy)

			for i := 0; i < xlimit; i++ {
				for j := 0; j < ylimit; j++ {
					pixel = getPixel(width, height, data, sx+i, sy+j)
					if pixel != targetPixelValue {
						ch |= dotPoisitions[i+(j*2)]
					}
				}
			}
			fmt.Printf("%c", (ch + 0x2800))
		}
		fmt.Print("\n")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
