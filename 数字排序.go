package main

import "fmt"

func main() {
	var (
		x int
		y int
		z int
	)
	fmt.Scanf("%d %d %d", &x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)

}
