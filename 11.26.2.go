package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func test1() {
	r, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	//r := strings.NewReader("hello world")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')

	fmt.Printf("s: %v\n", s)

}

func tset2() {
	r := strings.NewReader("ABCEFG")
	str := strings.NewReader("123456")
	r2 := bufio.NewReader(r)
	s, err := r2.ReadString('\n')
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("s: %v", s)
	r2.Reset(str)
	s2, err2 := r2.ReadString('\n')
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	fmt.Printf("s2: %v", s2)
}

func main() {
	tset2()
}
