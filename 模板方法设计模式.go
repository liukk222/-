package main

import "fmt"

type manyisen interface {
	cs()
	sh()
	sw()
	ys()
}
type man struct {
}

func (a *man) cs() {
	fmt.Println("出生")
}
func (a *man) sh() {
	fmt.Println("man")
}
func (a *man) sw() {
	fmt.Println("死亡")
}
func (a *man) ys() {
	a.cs()
	a.sh()
	a.sw()
}

type Cgman struct {
	man
}

func (b *Cgman) sh() {
	fmt.Println("成功的man")
}
func (b *Cgman) ys() {
	b.cs()
	b.sh()
	b.sw()
}

type Sbman struct {
	man
}

func (c *Sbman) sh() {
	fmt.Println("失败的man")
}
func (c *Sbman) ys() {
	c.cs()
	c.sh()
	c.sw()
}

func main() {
	s1 := man{}
	s2 := Cgman{
		s1,
	}
	s2.ys()
	fmt.Println("----------")
	s3 := Sbman{
		s1,
	}
	s3.ys()
}
