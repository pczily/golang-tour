package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	dx int
	dy int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.dx, img.dy)
}

func (img Image) At(x, y int) color.Color {

	var mpic = make([]([]uint8), img.dy)

	for row := 0; row < img.dy; row++ {
		mpic[row] = make([]uint8, img.dx)
		for col := 0; col < img.dx; col++ {
			mpic[row][col] = uint8((row + col))
		}
	}

	return color.RGBA{mpic[y][x], mpic[y][x], 255, 255}
}

func main() {
	m := Image{128, 128}
	pic.ShowImage(m)
}

