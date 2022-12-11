package main

import "fmt"

type Dz interface {
	dazhe() float32
}
type jiage struct {
	jg float32
}

func slja(a float32) *jiage {
	return &jiage{
		a,
	}
}

type DZ85 struct {
}

func Newdz85() *DZ85 {
	return &DZ85{}
}
func (*DZ85) dazhe() float32 {
	return 0.85
}
func (a *jiage) dzwan(d Dz) float32 {
	return a.jg * d.dazhe()
}

func main() {
	a := slja(100)
	b := Newdz85()
	fmt.Printf("a.dzwan(b): %v\n", a.dzwan(b))
}
