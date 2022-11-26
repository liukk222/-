package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func tset9() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	w := bufio.NewWriter(f)
	w.WriteString("hello,world!")
	w.Flush()
}
func test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	r := strings.NewReader("hello 世界")
	w := bufio.NewWriter(b)
	w.ReadFrom(r)
	fmt.Printf("b: %v\n", b)
}

func main() {
	//tset9()
	test10()
}
