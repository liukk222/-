package main

import "fmt"

// 使用字符串切片存储数据
type Stack []string

// 判断栈
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 入栈
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// 出栈
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index] // 获得顶部元素
		*s = (*s)[:index]      // 删除
		return element, true   // 返回
	}
}

func main() {
	var stack Stack // 创建一个栈

	stack.Push("java")
	stack.Push("python")
	stack.Push("golang!!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
