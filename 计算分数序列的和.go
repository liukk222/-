package main

import "fmt"

func main() {
	a := 1.0
	b := 2.0
	sun := 0.0
	for i := 0; i < 20; i++ {
		sun = sun + (b / a)
		b, a = a+b, b
	}
	fmt.Printf("sun: %f\n", sun)
}
