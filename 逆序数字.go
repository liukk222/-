package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Println("请输入一个五位以内的数：")
	fmt.Scanf("%d", &a)
	for i := 5; i > 0; i-- {
		x := a / int(math.Pow10(i-1))
		if x > 0 {
			fmt.Printf("数为%d,位数为%d\n", a, i)
			test(a, i)
			break
		}
	}

}
func test(x, y int) {
	if y > 1 {
		test(x, y-1)
	}
	r := x % int(math.Pow10(y)) / int(math.Pow10(y-1))
	fmt.Printf("%d", r)
}
