package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var rshu, monye int
	rshu = 10
	monye = 500
	for i := 0; i < 10; i++ {
		test1(rshu, monye)
		fmt.Println("")
	}
	s := 55 + 87 + 70 + 18 + 54 + 16 + 94 + 34 + 54 + 18
	fmt.Printf("s: %v\n", s)
}

func test1(Rshu, Monye int) {
	for i := 0; i < Rshu; i++ {
		m := test2(Rshu-i, Monye)
		fmt.Printf("%d ", m)
		Monye -= m
	}
}

func test2(Rshu, Monye int) int {
	if Rshu == 1 {
		return Monye
	}
	var min int = 1
	max := Monye / Rshu * 2
	rand.Seed(time.Now().UnixNano())
	Monye = rand.Intn(max) + min
	return Monye
}
