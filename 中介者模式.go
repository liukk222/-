package main

import (
	"fmt"
	"time"
)

// 联合国
type UnitedNations struct {
}

var un = NewUnitedNations()

// 创建联合国
func NewUnitedNations() *UnitedNations {
	return &UnitedNations{}
}

// 联合国显示各个国家发来的消息
func (un *UnitedNations) ShowMessage(nation *Nation, msg string) {
	fmt.Printf("%s: [ %s ]: %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		nation.Name,
		msg)
}

// 其他国家
type Nation struct {
	Name string
}

// new
func NewNation(name string) *Nation {
	return &Nation{name}
}

// SendMessage 各个国家向联合国发消息
func (n *Nation) SendMessage(msg string) {
	un.ShowMessage(n, msg)
}

func main() {

	n1 := NewNation("中国")
	n1.SendMessage("我要强大...")

	n2 := NewNation("美国")
	n2.SendMessage("我要看看谁不服...")
}
