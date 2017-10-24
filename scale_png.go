
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

/*
func blurSegment(seg image.Image) image.Image {
    return seg
}
*/

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
    pixelWidth := 25

    outImg := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
    for pixelX := bounds.Min.X; pixelX < bounds.Max.X; pixelX+= pixelWidth {
        for pixelY := bounds.Min.Y; pixelY < bounds.Max.Y; pixelY+= pixelWidth {

/*
            blurSegment(img.SubImage(image.Rectangle(pixelX, pixelY,
                                                     pixelX + pixelWidth,
                                                     pixelY + pixelWidth)))
*/

            r, g, b := 0, 0, 0
            for x := pixelX; x - pixelX < pixelWidth; x++ {
                for y := pixelY; y - pixelY < pixelWidth; y++ {
                    pr, pg, pb, _ := img.At(x,y).RGBA()
                    r += int(pr)
                    g += int(pg)
                    b += int(pb)
                    outImg.Set(x, y, img.At(x, y))
                }
            }

            var pixColor color.RGBA
            pixColor.R = uint8((r / (pixelWidth * pixelWidth )) >> 8)
            pixColor.G = uint8((g / (pixelWidth * pixelWidth )) >> 8)
            pixColor.B = uint8((b / (pixelWidth * pixelWidth )) >> 8)
            pixColor.A = 255

            for x := pixelX; x - pixelX < pixelWidth; x++ {
                for y := pixelY; y - pixelY < pixelWidth; y++ {
                    outImg.Set(x, y, pixColor)
                }
            }
        }
    }

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    imgEncoder := new(png.Encoder)
    imgEncoder.CompressionLevel = png.NoCompression
    err = imgEncoder.Encode(file, outImg)
    check(err)
}
