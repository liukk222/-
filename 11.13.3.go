package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}
func (cat Cat) eat() {
	fmt.Println("dog eat...")
}

func (cat Cat) sleep() {
	fmt.Println("dog sleep...")
}

type Person struct {
}

func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	person.care(dog)
	person.care(cat)
}
