/* package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("hello,%s", name)
}

func test(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	//test("tom", sayHello)
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)
	ff = cal("-")
	r = ff(2,1)
	fmt.Printf("r: %v\n", r)
} */

/* package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func max(a int, b int) int {
	if a > b {
		return a
type f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)inttype f1 func(int,int)int	} else {
		retu
	}
}

func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = max
	r = ff(1, 2)
	fmt.Printf("r: %v\n", r)
}
*/

package main

import "fmt"

func tesd1() {
	name := "tom "
	age := "20"

	f1 := func() string {
		return name + age
	}
	msg := f1()
	fmt.Printf("msg: %v\n", msg)
}
func main() {
	tesd1()

	//自己调用
	/* r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2) */

	//fmt.Printf("r: %v\n", r)

	/* max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	r := max(1, 2)
	fmt.Printf("r: %v\n", r) */

}
