package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		test(a)
		fmt.Printf("a: %v\n", a)
	}
}

func test(a []int) {
	var i, j, tep int
	rand.Seed(time.Now().UnixNano())
	for i = len(a) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		tep = a[i]
		a[i] = a[j]
		a[j] = tep
	}
}
