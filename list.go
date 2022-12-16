package main

import (
	"container/list"
	"fmt"
)

func Loop(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("e.Value: %v\n", e.Value)
	}
}

func Remove(e1 string, list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		if (e.Value).(string) == e1 {
			list.Remove(e)
			break
		}
	}
}

func main() {

	linkedlist := list.New()
	// 尾部添加
	linkedlist.PushBack("java")
	linkedlist.PushBack("python")
	linkedlist.PushBack("golang")

	// 遍历
	Loop(linkedlist)

	// 删除
	Remove("java", linkedlist)
	fmt.Println("------------")
	// 遍历
	Loop(linkedlist)
}
