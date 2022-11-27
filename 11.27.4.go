package main

import (
	"bytes"
	"fmt"
	"io"
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

func testBuffer3() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(b1[0:n]): %v\n", string(b1[0:n]))
	}
}

func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)
}

func testBuffer2() {
	var b bytes.Buffer
	n, err := b.WriteString("hello")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b.Bytes()[0:n]: %v\n", string(b.Bytes()[0:n]))
}

func testBuffer4() {
	var b = bytes.NewBufferString("hello wolrd")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("b1: %v\n", string(b1)[0:n])
	}

}

func main() {
	//testRepeat()
	//testCount()
	//testJoin()
	//testRunes()
	//testBuffer3()
	//testBuffer()
	//testBuffer2()
	testBuffer4()
}
