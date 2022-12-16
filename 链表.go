package main

import "fmt"

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域(指向下一个表的地址)
}

type List struct {
	headNode *Node //头节点
}

// 判断列表是否为空，如果头节点为空，则列表为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

// 获取列表长度
func (this *List) GetLenth() int {
	current := this.headNode
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	return len
}

// 从链表头部添加元素
func (this *List) Add(d Object) *Node {
	node := &Node{Data: d}
	node.Next = this.headNode
	this.headNode = node
	return node
}

// 从链表尾部添加元素
func (this *List) Append(d Object) {
	node := &Node{Data: d}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		current := this.headNode
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// 在链表指定位置添加元素
func (this *List) Insert(index int, data Object) {
	if index < 0 {
		this.Add(data)
	} else if index > this.GetLenth() {
		this.Append(data)
	} else {
		pre := this.headNode
		len := 0
		for len < (index - 1) {
			pre = pre.Next
			len++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除链表指定值的元素
func (this *List) Remove(data Object) {
	pre := this.headNode
	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// 删除指定位置的元素
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	if index <= 0 {
		this.headNode = pre.Next
	} else if index > this.GetLenth() {
		fmt.Println("超长了...")
		return
	} else {
		len := 0
		for len != (index-1) && pre.Next != nil {
			len++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// 查看链表是否包含某个元素
func (this *List) Contains(data Object) bool {
	current := this.headNode
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// 遍历链表所有节点
func (this *List) LoopList() {
	if !this.IsEmpty() {
		current := this.headNode
		for {
			fmt.Printf("current.Data: %v\n", current.Data)
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
	}
}

func main() {
	list := List{}
	list.Append("Java")
	list.Append("Python")
	list.Append("Golang")

	fmt.Printf("list.GetLenth(): %v\n", list.GetLenth())
	fmt.Println("-----------")
	list.LoopList()

	list.Add("before")
	fmt.Println("---------")

	list.LoopList()

	list.Remove("Java")

	fmt.Println("---------")
	list.LoopList()

	list.RemoveAtIndex(1)
	fmt.Println("----------")
	list.LoopList()
}
