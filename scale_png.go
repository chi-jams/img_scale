
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "bytes"
    "image/png"
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

    img, err := png.Decode(bytes.NewReader(dat))

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    img_encoder := new(png.Encoder)
    img_encoder.CompressionLevel = png.NoCompression
    err = img_encoder.Encode(file, img)
    check(err)
}
