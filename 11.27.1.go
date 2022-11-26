package main

import (
	"bufio"
	"fmt"
	"strings"
)

func test1() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		fmt.Printf("bs.Text(): %v\n", bs.Text())

	}
}

func main() {

}
