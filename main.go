package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var imagePath string

func init() {
    const (
        defaultPath = ""
        usage = "Path to image to be converted"
    )

    flag.StringVar(&imagePath, "image", defaultPath, usage)
    flag.StringVar(&imagePath, "i", defaultPath, usage+" (shorthand)")
}

func main() {
   flag.Parse()
   _, err := getImageFromPath(imagePath)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println("Image returned" ) 
}


func getImageFromPath(filePath string) (image.Image, error)  {
    f, err := os.Open(filePath) 
    defer f.Close()
    if err != nil {
        return nil, err
    }
    image, _, err := image.Decode(f)
    return image, err
}
