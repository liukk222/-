package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 10000; i++ {
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 268)))
		if x*x == (i+100) && y*y == (i+268) {
			fmt.Printf("i: %v\n", i)
		}
	}
}
