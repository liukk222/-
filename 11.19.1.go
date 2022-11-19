package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func makeDir() {
	/* err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func remove() {
	/* err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}

func wd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("dir: %v\n", dir)
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}
func reaFile() {
	b, err := os.ReadFile("tset.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b: %v\n", string(b[:]))
}

func writeFile() {
	os.WriteFile("tset.txt", []byte("你好，世界！"), os.ModePerm)
}

func main() {
	//createFile()
	//makeDir()
	//remove()
	//wd()
	//reaFile()
	writeFile()
}
