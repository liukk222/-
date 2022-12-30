package main

import "fmt"

type DaYing interface {
	daying()
}

type WenZi struct {
	wenzi string
}

func Wenzi(b WenZi) *WenZi {
	return &b
}
func (a *WenZi) daying() {
	fmt.Printf("%s", *a)
}

type JiaCu struct {
	jiacu DaYing
}

func Newjiacu(b DaYing) *JiaCu {
	return &JiaCu{
		jiacu: b,
	}
}
func (C *JiaCu) test() {
	fmt.Println("加粗")
}

func (d *JiaCu) test1() {
	d.jiacu.daying()
	d.test()
}

func main() {
	a := WenZi{
		wenzi: "hello,word",
	}
	b := Wenzi(a)
	c := Newjiacu(b)
	c.test1()
}
