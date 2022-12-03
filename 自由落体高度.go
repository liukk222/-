package main

import "fmt"

func main() {
	var mi float32 = 100.0
	s := mi / 2
	for i := 2; i <= 10; i++ {
		mi = mi + (s * 2)
		s = s / 2
	}
	fmt.Printf("s: %f\n", s)
	fmt.Printf("mi: %v\n", mi)
}
