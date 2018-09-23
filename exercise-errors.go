package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	x float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.x)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt{x}
		return 0, err
	} else {
		e := 1e-6
		z1 := 1.0
		z := 1.0
		z -= (z*z - x) / (2 * z)

		for (z-z1 > e) || (z1-z > e) {
			z1 = z
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
