package main

import (
    "github.com/fatih/color"
    "github.com/anthonynsimon/bild/effect"
    "github.com/anthonynsimon/bild/imgio"
    "github.com/anthonynsimon/bild/transform"
    "os"
    "fmt"
)

var (
  argchoose string
  filename string
  outputextension string
  width int
  height int
  radius float64
)

func main() {
    if len(os.Args) != 4 {
      usage()
      help()
      os.Exit(99)
    } else {
      filename = os.Args[1]
      argchoose = os.Args[2]
      outputextension = os.Args[3]

      img, err := imgio.Open(filename)
      if err != nil {
          println("The system can't find the specified file.")
          os.Exit(99)
      }

      effectchoose := effect.Invert(img)

      if argchoose == "Sobel" {
        effectchoose = effect.Sobel(img)
      } else if argchoose == "Invert" {
        effectchoose = effect.Invert(img)
      } else if argchoose == "Sepia" {
        effectchoose = effect.Sepia(img)
      } else if argchoose == "Sharpen" {
        effectchoose = effect.Sharpen(img)
      } else if argchoose == "EdgeDetection" {
        var radius float64
        fmt.Printf("Radius(float64) : ")
        fmt.Scanln(&radius)
        effectchoose = effect.EdgeDetection(img, radius)
      } else if argchoose == "Dilate" {
        var radius float64
        fmt.Printf("Radius(float64) : ")
        fmt.Scanln(&radius)
        effectchoose = effect.Dilate(img, radius)
      } else if argchoose == "Erode" {
        var radius float64
        fmt.Printf("Radius(float64) : ")
        fmt.Scanln(&radius)
        effectchoose = effect.Erode(img, radius)
      } else if argchoose == "Median" {
        var radius float64
        fmt.Printf("Radius(float64) : ")
        fmt.Scanln(&radius)
        effectchoose = effect.Median(img, radius)
      } else if argchoose == "UnsharpMask" {
        var (
          radius float64
          amount float64
        )
        fmt.Printf("Radius(float64) : ")
        fmt.Scanln(&radius)
        fmt.Printf("Amount(float64) : ")
        fmt.Scanln(&amount)
        effectchoose = effect.UnsharpMask(img, radius, amount)
      } else if argchoose == "Emboss" {
        effectchoose = effect.Emboss(img)
      } else {
        println("Bad [effect] Choose")
        os.Exit(99)
      }

      Cyanw := color.New(color.FgCyan, color.Bold)
      Redw := color.New(color.FgRed, color.Bold)
      Cyanw.Printf("\n\n\t\tImage Width - Height\n")
      Redw.Printf("\t\tWidth : ")
      fmt.Scanln(&width)
      Redw.Printf("\t\tHeight : ")
      fmt.Scanln(&height)
      Redw.Printf("\t\tRadius : ")
      fmt.Scanln(&radius)


      resized := transform.Resize(effectchoose, width, height, transform.Linear)
      rotated := transform.Rotate(resized, radius, nil)

      if outputextension == "jpeg" {
        if err := imgio.Save(filename[:len(filename)-3] + "jpeg", rotated, imgio.JPEGEncoder(95)); err != nil {
            panic(err)
        }
      } else if outputextension == "png" {
        if err := imgio.Save(filename[:len(filename)-3] + "png", rotated, imgio.PNGEncoder()); err != nil {
            panic(err)
        }
      } else if outputextension == "bmp" {
        if err := imgio.Save(filename[:len(filename)-3] + "bmp", rotated, imgio.BMPEncoder()); err != nil {
            panic(err)
        }
      } else {
        println("Bad output extension type")
        os.Exit(99)
      }
    }
}

func help() {
  Cyan := color.New(color.FgCyan, color.Bold)
  Red := color.New(color.FgRed, color.Bold)
  Cyan.Printf("\nEffects : \n")
  Red.Printf("\tSobel\n")
  Red.Printf("\tInvert\n")
  Red.Printf("\tSepia\n")
  Red.Printf("\tSharpen\n")
  Red.Printf("\tEdgeDetection\n")
  Red.Printf("\tDilate\n")
  Red.Printf("\tErode\n")
  Red.Printf("\tMedian\n")
  Red.Printf("\tUnsharpMask\n")
  Red.Printf("\tEmboss\n")
  Cyan.Printf("\nOutput Extensions : \n")
  Red.Printf("\tjpeg\n")
  Red.Printf("\tpng\n")
  Red.Printf("\tbmp\n\n")
}

func usage(){
  white := color.New(color.FgWhite)
  boldWhite := white.Add(color.Bold)
  boldWhite.Println("Usage : go run image-processor.go [filename] [effect] [outputextension]")
}
