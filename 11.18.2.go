package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var a int32 = 100

func add() {
	atomic.AddInt32(&a, 1)
}
func sub() {
	atomic.AddInt32(&a, -1)
}

/* var a = 100
var lock sync.Mutex

func add() {
	lock.Lock()
	a++
	lock.Unlock()
}

func sub() {
	lock.Lock()
	a--
	lock.Unlock()
}
func as() {
	lock.Lock()
	a = a * 2
	lock.Unlock()
}
func b() {
	lock.Lock()
	a = a / 2
	lock.Unlock()
}*/

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
		//go a()
		//go b()

	}
	time.Sleep(time.Second * 3)
	fmt.Printf("a: %v\n", a)

}
