package main

import "fmt"

func test1() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{4, 5, 6}
	i2 := append(s, s1...)
	fmt.Printf("i2: %v\n", i2)
	/* var s1 = make([]int, 2, 3)
	s1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("s1: %v\n", s1)
	*/
	//s1 = append(s1, 3)
	//fmt.Printf("s1: %v\n", s1)
	//s1 = append(s1, 3)
	/* var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:]
	fmt.Printf("s2: %v\n", s2) */
}

func testLen() {
	s := "hello world"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func testPrint() {
	name := "tom"
	age := 20
	print(name, " ", age, "\n")
	fmt.Println("------------")
	println(name, " ", age)
	fmt.Println("------------")
}

func main() {
	testPrint()
}
