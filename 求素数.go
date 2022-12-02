package main

import "fmt"

func test(x int) bool {
	for j := 2; j < x; j++ {
		if x%j == 0 {
			return false
		}
	}
	return true

}

func main() {
	b := 0
	for i := 1; i < 1000; i++ {
		a := test(i)
		if a {
			b++
			fmt.Printf("i: %v\n", i)
		}
	}
	fmt.Printf("b: %v\n", b)
}
