package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 0)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		defer close(chanInt)
		defer close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("Int: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("Str:%v\n", r)
		default:
			fmt.Println("defualt...")
		}
		time.Sleep(time.Second)
	}
}
