package main

import "fmt"

func test1() {
	var a1 = [2][2]int{{1, 2}, {3, 4}}

	fmt.Printf("a1: %v\n", a1)
	var a [3][3]string
	count := 3
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			a[i][j] = "*"
			print(a[i][j])
		}
		fmt.Println("\n")
	}
	add1 := make([][]int, count)
	for i := range add1 {
		add1[i] = make([]int, count)
	}
	fmt.Println(add1)
}

func main() {
	test1()

}
