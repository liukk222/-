package main

import (
	"fmt"
	"time"
)

func main() {
	/* timer := time.NewTicker(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-timer.C
	fmt.Printf("t1: %v\n", t1) */
	/* fmt.Printf("time.Now(): %v\n", time.Now())
	timer := time.NewTicker(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now()) */
	//time.Sleep(time.Second*2)
	/* fmt.Printf("time.Now(): %v\n", time.Now())
	<-time.After(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now()) */
	/* timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		fmt.Println("func...")
	}()
	s := timer.Stop()
	if s {
		fmt.Println("stop...")
	}

	time.Sleep(time.Second * 3) */
	fmt.Println("main ")
	timer := time.NewTicker(time.Second * 5)
	timer.Reset(time.Second * 1)
	<-timer.C
	fmt.Println("main end...")
}
