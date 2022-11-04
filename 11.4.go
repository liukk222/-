/*
	package main

import "fmt"

	func main() {
		s:="hello"
		fmt.Println("hello")
		fmt.Printf("s: %v\n", s)

fmt.Printf("\"hello\": %v\n", "hello")
}
*/
package main

import "fmt"

//import "c"

/* type WebSite struct {
	Name string
} */
//var a="ssr"
func main() {
	/* i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)
	i = 48
	fmt.Printf("i: %c\n", i)
	fmt.Printf("a:%x", 100) */
	/* x := 100
	p := &x
	fmt.Printf("p: %p\n", p) */

	/*  site := WebSite{Name: "duoke360"}
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("\"%%\": %v\n", "%%") */
	/* a := 10
	b := 5
	r := a == b
	fmt.Printf("r: %v\n", r)
	r = a > b
	fmt.Printf("r: %v\n", r)
	r = a != b
	fmt.Printf("r: %v\n", r)
	fmt.Printf("(a != b): %v\n", (a != b)) */
	/* var a [3][3]string
	count := 3
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			a[i][j] = "*"
			print(a[i][j])
		}
		print("\n")
	} */
	//fmt.Printf("a: %v\n", a)
	/* const PI float32 = 3.14
	const PI2 = 3.1415
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2) */
	/* const (
		a        = 100
		b        = 200
		c int    = 300
		d string = "liukk22"
	)
	fmt.Printf("d: %v\n", d) */
	a := 4
	b := 8
	fmt.Printf("a: %b\n", a)
	fmt.Printf("b: %b\n", b)
	fmt.Printf("(a & b): %v\n", (a & b))
	fmt.Printf("(a | b): %v\n", (a | b))
	r := a ^ b
	fmt.Printf("r: %v\n", r)
	r = a << 2
	fmt.Printf("r: %v\n", r)
	r = b >> 2
	fmt.Printf("r: %v\n", r)
}
