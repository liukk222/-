package main

import "fmt"

type Fruit interface {
	grant()
	pick()
}
type Apple struct {
	apple string
}
type Orang struct {
	orang string
}

func (*Apple) grant() {
	fmt.Println("种植苹果")
}
func (*Apple) pick() {
	fmt.Println("采栽苹果")
}
func (*Orang) pick() {
	fmt.Println("采栽橘子")
}
func (*Orang) grant() {
	fmt.Println("种植橘子")
}

func NewFruit(t int) Fruit {
	switch t {
	case 1:
		return &Apple{}

	case 2:
		return &Orang{}
	}
	return nil
}

func main() {
	apple := NewFruit(1)
	apple.grant()
	apple.pick()
	orang := NewFruit(2)
	orang.grant()
	orang.pick()
}
