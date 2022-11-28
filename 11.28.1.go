package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func mysearchInt(a *[]int, i int) {
	//i := len(*a)
	b := 1
	c := 0
	var f bool = true
	//w := i + 1
	for f {

		if (*a)[b] > (*a)[c] {

			fmt.Println("if")
			fmt.Printf("b: %d\n", b)
			fmt.Printf("c: %v\n", c)
			if b == i-1 {
				f = false
				break
			}
			b++
			c++
			//fmt.Printf("sb: %v\n", b)
		} else {
			Swap(&(*a)[b], &(*a)[c])
			fmt.Println("else")
			fmt.Printf("b: %d\n", b)
			fmt.Printf("c: %v\n", c)
			if b == i-1 {
				fmt.Println("ok")
				f = false
				myelse(a, i)
				break

			}
			b++
			c++

		}

	}

}

func myelse(a *[]int, i int) {
	//i := len(*a)
	bs := i - 2
	cs := i - 3

	var f bool = true
	for f {
		if (*a)[bs] > (*a)[cs] {
			if cs == 0 {
				f = false
				break
			}
			bs--
			cs--
		} else {
			Swap(&(*a)[bs], &(*a)[cs])
			fmt.Printf("cs: %v ", cs)
			fmt.Printf("bs: %v\n", bs)

			if cs == 0 {
				f = false
				break
			}
			bs--
			cs--

		}
	} //cs: 2 bs: 3
	//cs: 1 bs: 2
	//cs: 0 bs: 1//43215
	mysearchInt(a, i)
}

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	/* var f bool = true
	if f {
		fmt.Println("ok")
	} else {
		fmt.Println("aaa")
	} */
	s := []int{5, 4, 3, 2, 1}
	i := len(s)
	fmt.Printf("i: %d\n", i)
	mysearchInt(&s, i)
	fmt.Printf("s: %v\n", s)
	/* n := []uint{5, 4, 3, 2, 1}
	sort.Sort(NewInts(n))
	fmt.Println(n) */
}

//cs: 2 bs: 3
//cs: 1 bs: 2
//cs: 0 bs: 1//43215
