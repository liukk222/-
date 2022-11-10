package main

import "fmt"

func f2(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * f2(a-1)
	}
}

func f1() {
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func f3(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f3(n-1) + f3(n-2)
}

func main() {
	//f1()
	/* a := f2(5)
	fmt.Printf("a: %v\n", a) */
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")

}
