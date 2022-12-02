package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d int = 0, 0, 0, 0
	reader := bufio.NewReader(os.Stdin)
	ab, _ := reader.ReadString('\n')
	for _, v := range ab {
		switch {
		case v >= 'A' && v <= 'Z':
			a++
		case v >= 'a' && v <= 'z':
			a++

		case v >= ' ' || v <= '\t':
			b++
		case v >= '0' && v <= '9':
			c++
		default:
			d++

		}
	}
	d--
	fmt.Printf("char count = %d,space count = %d,digit count = %d,others count =%d\n", a, b, c, d)

}
