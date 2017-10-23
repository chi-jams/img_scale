
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
    "image"
    "image/color"
    "image/png"
    "image/jpeg"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %v <in_img> <out_img>\n", os.Args[0])
        os.Exit(1)
    }

    dat, err := ioutil.ReadFile(os.Args[1])
    check(err)

    img, err := jpeg.Decode(bytes.NewReader(dat))
    check(err)

    bounds := img.Bounds()
    //pixelWidth := 100
    outImg := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
    for pixelX := bounds.Min.X; pixelX < bounds.Max.X; pixelX++ {
        for pixelY := bounds.Min.Y; pixelY < bounds.Max.Y; pixelY++ {
            pr, pg, pb, pa := img.At(pixelX,pixelY).RGBA()
            var pixColor color.RGBA
            pixColor.R = uint8(pr>>8)
            pixColor.G = uint8(pg>>8)
            pixColor.B = uint8(pb>>8)
            pixColor.A = uint8(pa>>8)
            outImg.Set(pixelX, pixelY, pixColor)
        }
    }

/*
    for pixelX := bounds.Min.X; pixelX < bounds.Max.X; i+= pixelWidth {
        for pixelY := bounds.Min.Y; pixelY < bounds.Max.Y; j+= pixelWidth {
            pixelValue := 0
            for x := pixelX; x - pixelX < pixelWidth; x++ {
                for y := pixelY; y - pixelY < pixelWidth; y++ {
                    pixelValue += img.At(x, y)
                }
            }
            pixelValue /= pixelWidth ** 2
            for x := pixelX; x - pixelX < pixelWidth; x++ {
                for y := pixelY; y - pixelY < pixelWidth; y++ {
                    //outImg.Set(x, y, pixelValue)
                    pr, pg, pb, pa := img.At(x,y).RGBA()
                    pr := uint8(pr >> 8)
                    outImg.Set(x, y, img.At(i/pixelWidth * pixelWidth, j/pixelWidth * pixelWidth))
                }
            }
        }
    }
*/

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    imgEncoder := new(png.Encoder)
    imgEncoder.CompressionLevel = png.NoCompression
    err = imgEncoder.Encode(file, outImg)
    check(err)
}
