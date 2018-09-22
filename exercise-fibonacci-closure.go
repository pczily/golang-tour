package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fa, fb := 0, 1
	f := func() int {
		fc := fa + fb
		fShow := fa  // fShow for display
		fa = fb
		fb = fc
		return fShow
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
