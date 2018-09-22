package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var mpic = make([]([]uint8), dy)

	for row := 0; row < dy; row++ {
		mpic[row] = make([]uint8, dx)
		for col := 0; col < dx; col++ {
			mpic[row][col] = uint8(row * col)
		}
	}

	return mpic

}

func main() {
	pic.Show(Pic)
}
