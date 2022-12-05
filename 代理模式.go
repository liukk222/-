package main

import "fmt"

type Dadaima interface {
	Do()
}
type StrDdm struct {
}

func (a *StrDdm) Do() {
	fmt.Println("打代码")
}

type Newstrddm struct {
	StrDdm
}

func (b *Newstrddm) Do() {
	fmt.Println("打代码之前")
	b.StrDdm.Do()
	fmt.Println("打代码之后")
}

func main() {
	s1 := StrDdm{}
	s2 := Newstrddm{
		s1,
	}
	s2.Do()
}
