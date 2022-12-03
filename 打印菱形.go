package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 2-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= 2*i; k++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	for i := 0; i <= 2; i++ {
		for x := 0; x <= i; x++ {
			fmt.Printf(" ")
		}
		for y := 0; y <= (4 - 2*i); y++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

}
