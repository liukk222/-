// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func Swap(a *int, b *int) {
// 	tmp := *a
// 	*a = *b
// 	*b = tmp
// }

// func test(a *[]int) {
// 	//i := len(*a)
// 	//fmt.Printf("i: %v\n", i)
// 	x := 0
// 	v := 0
// 	for x = 0; x < len(*a); x++ {
// 		for v = 0; v < len(*a); v++ {
// 			if (*a)[x] < (*a)[v] {
// 				Swap(&(*a)[x], &(*a)[v])
// 			}

// 		}
// 	}
// 	//fmt.Printf("a: %v\n", a)
// }

// func test2() {
// 	arr := []int{1, 3, 5, 3, 4, 2, 1, 2, 3, 6, 8, 3}
// 	sort.Slice(arr, func(i, j int) bool {
// 		return arr[i] < arr[j]
// 	})
// 	fmt.Println(arr)
// }

// func mord(arr *[]int) {
// 	var a int = 0
// 	for i := 0; i < len(*arr); i++ {
// 		a = (*arr)[i]
// 		fmt.Printf("a: %v\n", a)
// 	}
// }

// type NewInts []uint

// func (n NewInts) Len() int {
// 	fmt.Printf("len: %v\n", len(n))
// 	return len(n)
// }

// func (n NewInts) Less(i, j int) bool {
// 	fmt.Println(i, j, n[i] < n[j], n)
// 	return n[i] < n[j]
// }

// func (n NewInts) Swap(i, j int) {
// 	n[i], n[j] = n[j], n[i]
// }

// func mysearchInt(a *[]int, i int) {
// 	//i := len(*a)
// 	b := 1
// 	c := 0
// 	//w := i + 1
// 	for {
// 		if (*a)[b] > (*a)[c] {
// 			b++
// 			c++
// 			fmt.Println("if")
// 			fmt.Printf("b: %d\n", b)
// 			if b == i {
// 				break
// 			}
// 		} else {
// 			Swap(&(*a)[b], &(*a)[c])
// 			b++
// 			c++
// 			fmt.Println("else")
// 			fmt.Printf("b: %d\n", b)
// 			if b > i {
// 				fmt.Println("ok")
// 				//myelse(a)
// 			}
// 		}

// 	}
// }

// func myelse(a *[]int) {
// 	i := len(*a)
// 	b := i - 2
// 	c := i - 3

// 	var f bool = false
// 	for f {
// 		if (*a)[b] > (*a)[c] {
// 			f = true
// 		} else {
// 			Swap(&(*a)[b], &(*a)[c])
// 			b--
// 			c--
// 			if c == 0 {
// 				break
// 			}
// 		}
// 	}
// 	//mysearchInt(&a)
// }

// func main() {
// 	//s := make([]int, 4, 8)
// 	s := []int{5, 4, 3, 2, 1} //4,3,2,1,5
// 	i := len(s)
// 	mysearchInt(&s, i)
// 	//var p *[5]int
// 	//p = &s
// 	//mord(p)
// 	//test(&s)
// 	//sort.Ints(s)
// 	//test(p)
// 	//test2()
// 	fmt.Printf("s: %v\n", s)
// 	//fmt.Printf("sort.SearchInts(s, 5): %v\n", sort.SearchInts(s, 5))
// 	//n := []uint{1, 2, 4, 3}
// 	//sort.Sort(NewInts(n))
// 	//fmt.Println(n)
// }
