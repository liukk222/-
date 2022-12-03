package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 1000; i++ {
		m := i
		for j := 1; j < i; j++ {
			if i%j == 0 {
				m = m - j
			}

		}

		if m == 0 {
			fmt.Printf("i: %v\n", i)
		}
	}
}
