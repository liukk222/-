package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	rand.Int63n(3)
	var a [3][3]string
	count := 3
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			a[i][j] = "*"
			print(a[i][j])
		}
		print("\n")
	}
	print("\n")
	g := 0
	for g <= 3 {
		q := rand.Int63n(3)
		w := rand.Int63n(3)
		if (q != 1) || (w != 1) && a[q][w] == "*" {
			a[q][w] = " "
			g++
		}
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			print(a[i][j])
		}
		print("\n")
	}
}
