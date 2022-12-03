package main

import (
	"fmt"
	"sync"
)

type Sinleton interface {
	dosomething()
}

// 首字母小写 私有的 不能导出
type sinleton struct {
}

func (s *sinleton) dosomething() {
	fmt.Println("do som thing")
}

var (
	once     sync.Once
	instance *sinleton
)

func GetInstance() Sinleton {
	once.Do(
		func() {
			instance = &sinleton{}
		},
	)
	return instance
}

func main() {
	s1 := GetInstance()
	fmt.Printf("s1: %p\n", s1)
	s2 := GetInstance()
	fmt.Printf("s2: %p\n", s2)
	s1.dosomething()
	s2.dosomething()
}
