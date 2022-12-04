package main

import "fmt"

type DaYing interface {
	daying(neirong string, zihao int)
}

type Zhong struct{}
type Ying struct{}

func (a Zhong) daying(neirong string, zihao int) {
	fmt.Printf("%s.字号是%d\n", zihao, zihao)
}
func (b Ying) daying(neirong string, zihao int) {
	fmt.Printf("%s.字号是%d\n", neirong, zihao)
}

type Neirong struct {
	Neirong string
	Zihao   int
	Apl     DaYing
}

func Newneirong(neirong string, zihao int, apl DaYing) *Neirong {
	return &Neirong{
		Neirong: neirong,
		Zihao:   zihao,
		Apl:     apl,
	}
}
func (sc *Neirong) Dy() {
	sc.Apl.daying(sc.Neirong, sc.Zihao)
}

func main() {
	a := Newneirong("hllo,word", 1, Ying{})
	a.Dy()
}
