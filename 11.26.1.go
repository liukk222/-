package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func testReadAll() {
	//r := strings.NewReader("hello world!")
	f, err2 := os.Open("a.txt")
	defer f.Close()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func testReaDir() {
	fi, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, v := range fi {
		//fmt.Printf("v: %v\n", v)
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func testReadFile() {
	b, err := ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func tsetWriteFile() {
	ioutil.WriteFile("a.txt", []byte("hello world"), 0664)
}

func main() {
	//testReaDir()
	//testReadFile()
	tsetWriteFile()
}
