package main

import "fmt"

func main() {
	var a, b, c, e, sum int
	fmt.Println("请输入二个数")
	fmt.Scanf("%d %d", &a, &b)
	for c < b {
		e = e + a
		sum = sum + e
		a = a * 10
		c++
	}
	fmt.Printf("sum: %v\n", sum)
}
