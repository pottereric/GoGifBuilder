package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

func main() {
	writeImageFile("rect.gif", DrawRedRectangle())

	writeImageFile("line.gif", buildLine())

}

func buildLine() gif.GIF {
	var images []*image.Paletted
	var delays []int

	width := 1000
	height := 1000

	palette := GetBWRBGPalette()
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for x := 0; x < width; x++ {
		// use the SetColorIndex method to set individual pixels to a palette color
		img.SetColorIndex(x, 200, 2)
	}

	images = append(images, img)
	delays = append(delays, 0)

	animation := gif.GIF{
		Image: images,
		Delay: delays,
	}
	return animation
}

func DrawRedRectangle() gif.GIF {
	var images []*image.Paletted
	var delays []int

	width := 200
	height := 200

	palette := GetBWRBGPalette()
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	red_rect := image.Rect(30, 30, 70, 70)

	// draw a shape onto the image using the Draw method
	draw.Draw(img, red_rect, &image.Uniform{palette[2]}, image.ZP, draw.Src)

	images = append(images, img)
	delays = append(delays, 0)

	animation := gif.GIF{
		Image: images,
		Delay: delays,
	}
	return animation
}

func GetBWRBGPalette() []color.Color {
	return []color.Color{
		color.Black,
		color.White,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}
}

func writeImageFile(fileName string, contents gif.GIF) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = gif.EncodeAll(f, &contents)

	if err != nil {
		fmt.Println(err)
		return
	}
}
