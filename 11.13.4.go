package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Printf("per: %v\n", per)
	fmt.Println("eat...")
}

func (per Person) sllep() {
	fmt.Println("sllep...")
}

func (per Person) work() {
	fmt.Println("work...")
}

func main() {
	per := Person{
		name: "tom",
		age:  20,
	}
	fmt.Printf("per: %v\n", per)
	per.eat()
	per.sllep()
	per.work()
}
