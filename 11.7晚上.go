package main

import "fmt"

func tesd4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	_, ok := m1[k1]
	//fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("--------")
	_, ok = m1[k2]
	fmt.Printf("ok: %v\n", ok)
	//fmt.Printf("v: %v\n", v)
}

func tsed1() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
}

func tsed5() {
	var s1 = []int{1, 2, 3, 4}
	s2 := make([]int, 4)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func tese3() {
	var s1 = []int{1, 2, 3, 4, 5}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

func tese2() {
	var s1 = []int{1, 2, 3, 4, 5}
	s1 = append(s1[:1], s1[3:]...)

	fmt.Printf("s1: %v\n", s1)
}

func test1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}
func main() {
	//test1()
	//tese2()
	//tese3()
	//tsed5()
	//tsed1()
	tesd4()
}

// package main

// import "fmt"

// func main() {
// 	var arr = [3]int{1, 2, 3}
// 	s5 := arr[:]
// 	fmt.Printf("arr: %d\n", arr[1])
// 	fmt.Printf("s5: %v\n", s5)
// }
