package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}

		}
		fmt.Println("")
	}
}
