package main

import "fmt"

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// 升序
func BubbleAsort(values *[]int) {
	// 有几个数比较几次
	for i := 0; i < len(*values)-1; i++ {
		// 比较总长度，每次减一
		for j := i + 1; j < len(*values); j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if (*values)[i] > (*values)[j] {
				// 交换
				Swap(&(*values)[i], &(*values)[j])
			}
		}
	}
	fmt.Println(*values)
}

// 降序
func BubbleZsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			// 如果后面的比前面的大就交换，所以是降序
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}
func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	BubbleAsort(&values)

}
