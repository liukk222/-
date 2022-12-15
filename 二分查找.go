package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := 0
	c := len(a)
	d := 8
	for b < c {
		e := (b + c) >> 1
		if d < a[e] {
			c--
		} else if d > a[e] {
			b++
		} else {
			fmt.Printf("%d", e)
			break
		}

	}
}
