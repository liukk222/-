package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println("请输入一个数")
	fmt.Scanf("%d", &a)
	for i := 2; i <= a; i++ {
		for a != i {
			if a%i == 0 {
				fmt.Printf("%d*", i)
				a = a / i
			} else {
				break
			}
		}
	}
	fmt.Println(a)
}
