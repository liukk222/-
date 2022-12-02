package main

import "fmt"

func main() {
	var (
		a int = 1
		b int = 1
	)
	for i := 0; i < 12; i++ {
		fmt.Println(a, b)
		a += b
		b += a

	}
}
