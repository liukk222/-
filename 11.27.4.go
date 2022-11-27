package main

import (
	"bytes"
	"fmt"
	"strings"
)

func testTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)
	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
}

func testContais() {
	s := "duoke360.com"
	b := []byte(s)
	b1 := []byte("duoke360")
	b2 := []byte("Duoke360")
	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("strings.Contains(\"hello\", \"hello\"): %v\n", strings.Contains("hello", "hello"))
}

func testCount() {
	s := []byte("helloooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Printf("bytes.Count(s, sep1): %v\n", bytes.Count(s, sep1))
	fmt.Printf("bytes.Count(s, sep2): %v\n", bytes.Count(s, sep2))
	fmt.Printf("bytes.Count(s, sep3): %v\n", bytes.Count(s, sep3))
}

func testRepeat() {
	b := []byte("hi")
	fmt.Printf("string(bytes.Repeat(b, 1)): %v\n", string(bytes.Repeat(b, 1)))
	fmt.Printf("string(bytes.Repeat(b, 3)): %v\n", string(bytes.Repeat(b, 3)))
}
func testRunes() {
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("len(r): %v\n", len(r))
}
func testJoin() {
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4)))
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5)))
}

func main() {
	//testRepeat()
	//testCount()
	testJoin()
	//testRunes()
}
