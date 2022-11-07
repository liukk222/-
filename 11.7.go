package main

import "fmt"

func f1() {
	var a = [3]int{1, 2, 3}
	/* for i, v := range a {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	} */
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	f1()
}

/* package main

import "fmt"

func a1() {
	var a1 = [...]int{0: 1, 3: 100, 5: 10}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]bool{2: true, 5: false}
	fmt.Printf("a2: %v\n", a2)
}

func main(fmt
	//a1()
	var (
		a int = 1
		b int = 2
	)
	a=a+b
	b=a-b
	a=a-b
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
*/
