package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//fmt.Printf("math.Mod(20, 2): %v\n", math.Mod(20, 2))
	//fmt.Printf("math:", math.(10, 6))
	//fmt.Printf("math.Pi: %v\n", math.Pi)
	/* a := 10
	c := 2
	b := a % c
	fmt.Printf("b: %v\n", b) */

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("------------")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	fmt.Println("------------")
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
