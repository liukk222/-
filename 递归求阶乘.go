package main

import "fmt"

func test(i int) (sum int) {
	if i == 0 {
		return 1
	}
	sum += i * test(i-1)
	return
}

func main() {
	sum := test(5)
	fmt.Printf("sum: %v\n", sum)
}
