package main

import (
	"fmt"
)

func main() {

	fmt.Printf("请输入 %d 个字符: ", 5)
	putchar(5)
}
func putchar(n int) {
	var c byte
	if n >= 1 {
		fmt.Scanf("%c", &c) //abcde
		putchar(n - 1)
		fmt.Printf("%c", c)
	}
}
