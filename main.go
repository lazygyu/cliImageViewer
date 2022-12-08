package main

import (
  "fmt"
  "gv/dotscreen"
  "image"
  _ "image/jpeg"
  _ "image/png"
  _ "image/gif"
  "image/draw"
  "log"
  "os"
  "flag"

  "github.com/nfnt/resize"
  "golang.org/x/term"
)

func main() {

  invert := flag.Bool("i", false, "Invert the image (default: false)")
  var maxWidth int
  var maxHeight int

  flag.IntVar(&maxWidth, "w", 0, "Limit the width of the image (if the limitation is larger than the screen width, it will be ignored)")
  flag.IntVar(&maxHeight, "h", 0, "Limit the height of the image (if the limitation is larger than the screen height, it will be ignored)")

  flag.Parse()

  args := flag.Args()

  if len(args) < 1 {
    fmt.Println("Usage: gv [options] filename")
    flag.PrintDefaults()
    os.Exit(0)
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

  stdout := os.Stdout.Fd()

  if (term.IsTerminal(int(stdout))) {
    width, height, err := term.GetSize(int(stdout))
    if err != nil {
      log.Fatal("Cannot get the size of the terminal")
    }
    width *= 2
    height *= 5
    if maxWidth == 0 || width < maxWidth {
      maxWidth = width
    }
    if maxHeight == 0 || height < maxHeight {
      maxHeight = height
    }
  }

  if maxWidth > 0 && maxHeight > 0 {
    width := maxWidth
    height := maxHeight


    // resize to fit
    if img.Bounds().Max.X > width && img.Bounds().Max.Y > height {
      xratio := float32(img.Bounds().Max.X) / float32(width)
      yratio := float32(img.Bounds().Max.Y) / float32(height)
      var w, h int
      if xratio >= yratio {
        w = width
        h = int(float32(img.Bounds().Max.Y) / xratio)
      } else {
        w = int(float32(img.Bounds().Max.X) / yratio)
        h = height
      }
      img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
    } else if img.Bounds().Max.X > width {
      ratio := float32(img.Bounds().Max.X) / float32(width)
      w := width
      h := int(float32(img.Bounds().Max.Y) / ratio)
      img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
    } else if img.Bounds().Max.Y > height {
      ratio := float32(img.Bounds().Max.Y) / float32(height)
      w := int(float32(img.Bounds().Max.X) / ratio)
      h := height
      img = resize.Resize(uint(w), uint(h), img, resize.Lanczos2)
    }
  }

  nimg := toGrayscale(img)
  draw.Draw(nimg, nimg.Bounds(), img, img.Bounds().Min, draw.Src)

  // dithering
  result := Dithering(*nimg)
  // print out
  dotscreen.PrintImage(img.Bounds().Max.X, img.Bounds().Max.Y, &result, *invert)
}
