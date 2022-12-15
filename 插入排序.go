package main

import "fmt"

func main() {
	a := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	for i := 0; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				tem := a[j]
				a[j] = a[j-1]
				a[j-1] = tem
			}
		}
	}
	fmt.Printf("a: %v\n", a)
}
