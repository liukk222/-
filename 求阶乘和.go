package main

import "fmt"

func main() {
	a := 1
	sum := 0
	for i := 1; i <= 20; i++ {
		a = a * i
		sum = sum + a
	}
	fmt.Printf("sum: %v\n", sum)
}
