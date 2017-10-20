
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
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
    SCALE := 100
    pixel_width := (bounds.Max.Y - bounds.Min.Y) / SCALE

    out_img := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
    for i := bounds.Min.X; i < bounds.Max.X; i++ {
        for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
            out_img.Set(i, j, img.At(i/pixel_width * pixel_width, j/pixel_width * pixel_width))
        }
    }

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    img_encoder := new(png.Encoder)
    img_encoder.CompressionLevel = png.NoCompression
    err = img_encoder.Encode(file, out_img)
    check(err)
}
