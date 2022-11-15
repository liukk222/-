package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var i int = 100
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}
func sub() {
	lock.Lock()
	defer wg.Done()
	i -= 1
	fmt.Printf("i--: %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end i: %v\n", i)
}
