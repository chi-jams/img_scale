
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "image/jpeg"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func blurSquare(img *image.RGBA, pixelX, pixelY, pixelWidth int) {
    r, g, b := 0, 0, 0
    for x := pixelX; x - pixelX < pixelWidth; x++ {
        for y := pixelY; y - pixelY < pixelWidth; y++ {
            pr, pg, pb, _ := img.At(x,y).RGBA()
            r += int(pr)
            g += int(pg)
            b += int(pb)
        }
    }

    var pixColor color.RGBA
    pixColor.R = uint8((r / (pixelWidth * pixelWidth )) >> 8)
    pixColor.G = uint8((g / (pixelWidth * pixelWidth )) >> 8)
    pixColor.B = uint8((b / (pixelWidth * pixelWidth )) >> 8)
    pixColor.A = 255

    draw.Draw(img, image.Rect(pixelX, pixelY, pixelX + pixelWidth, pixelY + pixelWidth), &image.Uniform{pixColor}, image.ZP, draw.Src)
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %v <in_img> <out_img>\n", os.Args[0])
        os.Exit(1)
    }

    dat, err := ioutil.ReadFile(os.Args[1])
    check(err)

    rawImg, err := jpeg.Decode(bytes.NewReader(dat))
    bounds := rawImg.Bounds()
    img := image.NewRGBA(rawImg.Bounds())
    draw.Draw(img, img.Bounds(), rawImg, rawImg.Bounds().Min, draw.Src)
    check(err)

    pixelWidth := 25
    for pixelX := bounds.Min.X; pixelX < bounds.Max.X; pixelX+= pixelWidth {
        for pixelY := bounds.Min.Y; pixelY < bounds.Max.Y; pixelY+= pixelWidth {
            go blurSquare(img, pixelX, pixelY, pixelWidth)
        }
    }

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    imgEncoder := new(png.Encoder)
    imgEncoder.CompressionLevel = png.NoCompression
    err = imgEncoder.Encode(file, img)
    check(err)
}
