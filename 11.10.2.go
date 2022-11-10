

/* package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var pa [3]*int
	fmt.Printf("pa: %v\n", pa)
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Printf("pa: %v\n", pa)
	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v\n", *pa[i])
	}
} */

// /* package main

// import "fmt"

// func main() {
// 	var ip *int
// 	fmt.Printf("ip: %v\n", ip)
// 	fmt.Printf("ip: %T\n", ip)
// 	var i int = 100
// 	ip = &i
// 	fmt.Printf("ip: %v\n", ip)
// 	fmt.Printf("ip: %v\n", *ip)
// } */

// package main

// import "fmt"

// var i int = initVar()

// func initVar() int {
// 	fmt.Println("initVar...")
// 	return 100
// }

// func init() {
// 	fmt.Println("init...")
// }
// func init() {

// }
// func main() {
// 	fmt.Println("main")
// }
