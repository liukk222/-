package main

import "fmt"

func main() {
	var a, b, c int16
	for a = 'x'; a <= 'z'; a++ {
		for b = 'x'; b <= 'z'; b++ {
			if a != b {
				for c = 'x'; c <= 'z'; c++ {
					if a != c && b != c {
						if c != 'x' && c != 'z' && a != 'x' {
							fmt.Printf("a--%c b--%c c--%c", a, b, c)
						}
					}
				}
			}
		}
	}
}
