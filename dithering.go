package main

import (
  "image"
  "image/draw"
)

func Dithering(img image.Gray) []byte {
  var oldPixel, newPixel int
  var error int
  width := img.Bounds().Max.X
  height := img.Bounds().Max.Y

  origin := toByteArray(img)
  result := make([]byte, width * height, width * height)

  var min, max byte
  var half int

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

  half = (int(min) + int(max)) / 2

  e := make([]int, width*2)
  m := []int{0, 1, width-2, width-1, width, 2*width-1}

  l := len(origin)

  for idx := 0; idx < l; idx++ {
    oldPixel = int(origin[idx]) + e[0]
    e = append(e[1:], 0)
    newPixel = findColor(oldPixel, half)
    error=(oldPixel - newPixel)>>3
    result[idx] = byte(newPixel / 255)
    for _, k := range m {
      e[k] += error
    }
  }
  return result
}

func addV(target *[]byte, pos, v int) {
  arr := *target
  if (pos >= len(arr)) {
    return
  }
  arr[pos] = byte(int(arr[pos]) + v)
}

func toByteArray(img image.Gray) []byte {
  width := img.Bounds().Dx()
  height := img.Bounds().Dy()

  origin := make([]byte, width*height)
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      origin[y*width+x] = img.GrayAt(x, y).Y
    }
  }
  return origin
}

func toGrayscale(img image.Image) *image.Gray {
  result := image.NewGray(img.Bounds())
  draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)
  return result
}

func findColor(c int, half int) int {
  if c > half {
    return 255
  } else {
    return 0
  }
}
