package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	sentence := strings.Fields(s)

	for _, v := range sentence {
		_, ok := count[v]
		if ok == true {
			count[v]++
		} else {
			count[v] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
