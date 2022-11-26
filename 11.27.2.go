package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

func test2() {
	defer fmt.Println("panic 结束后再执行...")
	log.Print("my log")
	log.Panic("my panic")
	fmt.Println("end...")
}

func test3() {
	defer fmt.Println("defer...")
	log.Print("my log")
	log.Fatal("fatal...")
	fmt.Println("end...")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	//defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.Print("my log")
}
