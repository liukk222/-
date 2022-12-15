package main

import "fmt"

func quickSort(arr []int) []int {
	// 总长度
	length := len(arr)
	// 判断长度是否为1，一个元素就不用排了
	if length <= 1 {
		return arr
	}
	// 假设第一个元素是中间值
	middle := arr[0]
	// 左边元素
	var left []int
	// 右边元素
	var right []int
	for i := 1; i < length; i++ {
		if middle < arr[i] {
			right = append(right, []int{arr[i]}...)
		} else {
			left = append(left, []int{arr[i]}...)
		}
	}

	// 递归排左边
	left = quickSort(left)
	// 递归排右边
	right = quickSort(right)
	middle_s := []int{middle}
	result := append(append(left, middle_s...), right...)

	return result
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	result := quickSort(values)
	fmt.Printf("result: %v\n", result)
}
