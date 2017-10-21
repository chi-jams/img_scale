
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
    "reflect"
    "image"
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

    bounds := img.Bounds()
    //pixelWidth := 100

    fmt.Println(img.At(0,0))
    fmt.Println(reflect.TypeOf(image.YCbCrToRGB(img.At(0,0))))
    outImg := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
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
                    outImg.Set(x, y, pixelValue)
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
