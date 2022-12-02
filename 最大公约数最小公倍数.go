package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Println("请输入")
	fmt.Scanf("%d %d", &a, &b)
	c = a * b
	for b != 0 {
		d = a % b
		a = b
		b = d
	}

	fmt.Printf("最大共约数是%d,最小共倍数是%d", a, (c / a))
}
