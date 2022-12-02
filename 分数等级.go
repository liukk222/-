package main

import "fmt"

type B bool

func (b B) C(t, f interface{}) interface{} {
	if bool(b) == true {
		return t
	}
	return f
}

func main() {
	fmt.Println("请输入一个数")
	var a int = 0
	fmt.Scanf("%d", &a)
	b := B(a >= 90).C("A", B(a >= 60).C("B", "C"))
	fmt.Printf("b: %v\n", b)
}
