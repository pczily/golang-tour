package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	e := 1e-6
	z1 := 1.0
	z := 1.0
	z -= (z*z - x) / (2 * z)

	for (z-z1 > e) || (z1-z > e) {
		z1 = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
