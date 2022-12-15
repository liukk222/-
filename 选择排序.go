package main

import "fmt"

func main() {
	var a = []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
			tem := a[i]
			a[i] = a[min]
			a[min] = tem
		}
	}
	fmt.Printf("a: %v\n", a)
}
