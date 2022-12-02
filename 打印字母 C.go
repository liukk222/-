package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i < 3 {
			for j := 4; j > 2*i-1; j-- {
				fmt.Print("*")
			}

		} else {
			for k := 1; k < 2*i-2; k++ {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}

}
