package main

import "fmt"

type Zimu interface {
	a()
	b()
	c()
}
type Xiaoxie struct {
	Name string
}

func (this *Xiaoxie) a() {
	fmt.Println("a")
}
func (this *Xiaoxie) b() {
	fmt.Println("b")
}
func (this *Xiaoxie) c() {
	fmt.Println("c")
}

type Daxie struct {
	Name string
}

func (this *Daxie) a() {
	fmt.Println("A")
}
func (this *Daxie) b() {
	fmt.Println("B")
}
func (this *Daxie) c() {
	fmt.Println("C")
}

type Diz struct {
	zimu Zimu
}

func NewConstruct(z Zimu) *Diz {
	return &Diz{
		zimu: z,
	}
}
func (this *Diz) daying() {
	this.zimu.a()
	this.zimu.b()
	this.zimu.c()
}

func main() {
	da := &Daxie{}
	d := NewConstruct(da)
	d.daying()
	fmt.Println("-----------")
	xiao := &Xiaoxie{}
	s := NewConstruct(xiao)
	s.daying()
}
