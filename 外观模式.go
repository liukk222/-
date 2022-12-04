package main

import "fmt"

type DaYing interface {
	daying()
}
type DaXie struct {
}
type XiaoXie struct{}

func testda() *DaXie {
	return &DaXie{}
}
func testxiao() *XiaoXie {
	return &XiaoXie{}
}
func (a *DaXie) daying() {
	fmt.Println("A")
}
func (b *XiaoXie) daying() {
	fmt.Println("a")
}

type jihe struct {
	Daxie   DaXie
	Xiaoxie XiaoXie
}

func testjh() *jihe {
	return &jihe{
		Daxie:   DaXie{},
		Xiaoxie: XiaoXie{},
	}
}

func (f *jihe) Newjh() {
	f.Daxie.daying()
	f.Xiaoxie.daying()
}

func main() {
	s1 := testjh()
	s1.Newjh()
}
