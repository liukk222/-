package main

import (
	"fmt"
)

func mergerSort(arr []int, a, b int) {
	// 长度小于等于1，不用排序
	if b-a <= 1 {
		return
	}
	// 取中值
	c := (a + b) / 2
	// 递归调用
	// 排左边
	mergerSort(arr, a, c)
	// 排右边
	mergerSort(arr, c, b)
	// 左边切片
	arrLeft := make([]int, c-a)
	// 右边切片
	arrRight := make([]int, b-c)
	fmt.Println("")
	// 拷贝左边数据到左边切片
	copy(arrLeft, arr[a:c])
	fmt.Printf("arrLeft: %v ", arrLeft)
	fmt.Printf("%d:%d\n", a, c)
	// 拷贝右边数据到右边切片
	copy(arrRight, arr[c:b])
	fmt.Printf("arrRight: %v ", arrRight)
	fmt.Printf("%d:%d\n", c, b)
	i := 0
	j := 0
	for k := a; k < b; k++ {
		if i >= c-a {
			arr[k] = arrRight[j]
			j++
			fmt.Println("c-a")
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
			fmt.Println("b-c")
		} else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
			fmt.Println("arrLeft[i] < arrRight[j]")
		} else {
			arr[k] = arrRight[j]
			j++
			fmt.Println("else")
		}
		fmt.Printf("%v\n", arr)
	}
}

func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(arr)
	mergerSort(arr, 0, len(arr))
	fmt.Println(arr)
}
