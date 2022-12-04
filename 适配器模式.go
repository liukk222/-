package main

import "fmt"

type DaYing interface {
	Daying()
}
type Daxie struct {
}

func (Daxie) Daying() {
	fmt.Println("A")
}

type XiaoXie struct {
	Daxie
}
type daying interface {
	dayings()
}

func (a *XiaoXie) dayings() {
	fmt.Println("a")
	a.Daying()
}

func main() {
	daxie := Daxie{}
	xiaoxie := XiaoXie{Daxie: daxie}
	xiaoxie.dayings()

}
