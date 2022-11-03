package main

import "fmt"

func getNamendAge() (string, int) {
	return "tom", 20
}

func main() {
	// var (
	// name string="tom"
	// age  int=12
	// b    bool=true
	// )
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// fmt.Printf("b: %v\n", b)
	// var name="tom"
	// var age =20
	// var b=true
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// fmt.Printf("b: %v\n", b)
	// name :="tom"
	// age:=20
	name, age := getNamendAge()
	b := true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)
	// asdsffsd
}
