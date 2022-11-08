package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		fmt.Printf("a1: %v\n", a)
		return base
	}
	sub := func(a int) int {
		base -= a
		fmt.Printf("a2: %v\n", a)
		return base
	}
	return add, sub
}

func main() {
	/* f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("---------")
	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(200)
	fmt.Printf("r: %v\n", r) */
	add, sub := cal(100)
	r := add(100)
	fmt.Printf("r: %v\n", r)
	r = add(30)
	fmt.Printf("r: %v\n", r)
	r = sub(50)
	fmt.Printf("r: %v\n", r)
	r = sub(30)
	fmt.Printf("r: %v\n", r)
	add1, sub1 := cal(100)
	j := add1(2)
	fmt.Printf("j: %v\n", j)
	j = sub1(1)
	fmt.Printf("j: %v\n", j)
}
