package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog struct {
	Animal
	//name string
	color string
}

type Cat struct {
	Animal
	color string
}

func main() {
	/* dog := Dog{
		a:     Animal{name: "花花", age: 2},
		color: "褐色",
	}
	dog.a.eat()
	dog.a.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)
	fmt.Printf("dog.a.age: %v\n", dog.a.age) */
	cat := Cat{
		Animal{name: "黑黑", age: 1},
		"白色",
	}
	cat.eat()
	cat.sleep()
	fmt.Printf("cat: %v\n", cat)
}
