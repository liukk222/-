package main

import "fmt"

func add(x int, y int) (ret int) {
	ret = x + y
	return ret
}

func comp(a int, b int) (max int) {
	if a > b {
		max = a
		return max
	} else {
		max = b
		return max
	}
}
func f1(a int) {
	a = 200
}
func f2(s []int) {
	s[0] = 1000
}
func f3(s ...int) {
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	f3(1, 2, 3, 4, 5, 6)
	f3(4, 5, 6, 7, 8)

	/* s := []int{1, 2, 3}
	f2(s)
	fmt.Printf("s: %v\n", s) */

	/* a := 12
	b := 14
	c := comp(a, b)
	fmt.Printf("comp(a, b): %v\n", comp(a, b))
	fmt.Printf("c: %v\n", c) */
	/* x := 200
	f1(x)
	fmt.Printf("x: %v\n", x) */
}

/* package main

import "fmt"

func tesd1() {
	var s1 = []int{1, 2, 3, 4, 5}
	l := len(s1)
	for i := 0; i < l; i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
	}
}

func tesd2() {
	var s1 = []int{1, 2, 3, 4, 5}
	for i, v := range s1 {
		fmt.Printf("s1[%v]: %v\n", i, v)
	}
}

func main() {
	//tesd1()
	tesd2()
}
*/
