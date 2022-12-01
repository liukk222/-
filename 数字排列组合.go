package main

import "fmt"

func main() {
	count := 5
	a := 0
	for i := 1; i < count; i++ {
		for j := 1; j < count; j++ {
			for k := 1; k < count; k++ {
				if i != j && i != k && k != j {
					a++
					fmt.Printf("第%d种方案,i=%d,j=%d,k=%d\n", a, i, j, k)
				}
			}
		}
	}
}
