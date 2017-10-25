
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
    "sync"
    "runtime"
)

// We're just gonna explode if something goes wrong
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Blurs a square of pixWidth at pixX, piY from rawImg to img
func pixSquare(img *image.RGBA, pixX, pixY, pixWidth int, rawImg *image.YCbCr) {
    r, g, b := 0, 0, 0
    for x := pixX; x - pixX < pixWidth; x++ {
        for y := pixY; y - pixY < pixWidth; y++ {
            pr, pg, pb, _ := rawImg.At(x,y).RGBA()
            r += int(pr)
            g += int(pg)
            b += int(pb)
        }
    }

    var pixColor color.RGBA
    pixColor.R = uint8((r / (pixWidth * pixWidth )) >> 8)
    pixColor.G = uint8((g / (pixWidth * pixWidth )) >> 8)
    pixColor.B = uint8((b / (pixWidth * pixWidth )) >> 8)
    pixColor.A = 255

    draw.Draw(img, image.Rect(pixX, pixY, pixX + pixWidth, pixY + pixWidth),
              &image.Uniform{pixColor}, image.ZP, draw.Src)
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %v <in_img> <out_img>\n", os.Args[0])
        os.Exit(1)
    }

    dat, err := ioutil.ReadFile(os.Args[1])
    check(err)

    blep, err := jpeg.Decode(bytes.NewReader(dat))
    rawImg := blep.(*image.YCbCr)
    check(err)

    img := image.NewRGBA(rawImg.Bounds())

    pixWidth := 25
    bounds := img.Bounds()
    var wg sync.WaitGroup
    stripHeight := bounds.Max.Y / runtime.NumCPU()
    for startStrip := 0; startStrip < bounds.Max.Y; startStrip += stripHeight {
        wg.Add(1)
        go func(img *image.RGBA, startStrip, stripHeight, pixWidth int, rawImg *image.YCbCr) {
            defer wg.Done()

            for pixX := bounds.Min.X; pixX < bounds.Max.X; pixX += pixWidth {
                for pixY := startStrip; pixY < startStrip + stripHeight; pixY += pixWidth {
                    pixSquare(img, pixX, pixY, pixWidth, rawImg)
                }
            }
        }(img, startStrip, stripHeight, pixWidth, rawImg)
    }

    wg.Wait()

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    imgEncoder := new(png.Encoder)
    imgEncoder.CompressionLevel = png.NoCompression
    err = imgEncoder.Encode(file, img)
    check(err)
}
