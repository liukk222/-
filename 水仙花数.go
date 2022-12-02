package main

import "fmt"

func main() {
	x := 0
	y := 0
	z := 0

	for i := 0; i < 1000; i++ {
		if i > 100 {
			x = i / 100
			y = i / 10 % 10
			z = i % 10
			if i == x*x*x+y*y*y+z*z*z {
				fmt.Printf("i: %v\n", i)
			}
		}
	}
}
