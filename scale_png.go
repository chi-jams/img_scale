
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

    fmt.Println("Still here baybee!")
    fmt.Println(os.Args)

    dat, err := ioutil.ReadFile(os.Args[1])
    check(err)

    img, err := png.Decode(bytes.NewReader(dat))
    fmt.Println(img.Bounds().Max.X)
    fmt.Println(img.Bounds().Max.Y)

    file, err := os.Create(os.Args[2])
    defer file.Close()
    check(err)

    img_encoder := new(png.Encoder)
    //img_encoder.CompressionLevel = png.NoCompression
    err = img_encoder.Encode(file, img)
    check(err)

    //err = ioutil.WriteFile(os.Args[2], dat, 0644)
    //check(err)
}
