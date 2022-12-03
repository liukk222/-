package main

import (
	"fmt"
	"math"
)

func main() {

	var a, i int
	fmt.Println("请输入一个五位以内的数：")
	fmt.Scanf("%d", &a)
	var max int = 5
	for i = 0; i < max/2; i++ {
		h := a % int(math.Pow10(max-i)) / int(math.Pow10(max-i-1))
		l := a % int(math.Pow10(i+1)) / int(math.Pow10(i))
		if h != l {
			fmt.Println(a, "不为回文数")
			break
		}
	}
	if i >= max/2 {
		fmt.Println(a, "是为回文数")
	}

}
