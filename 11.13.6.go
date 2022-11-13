package main

import "fmt"

type Person struct {
	name string
	age  int
}

func Newperson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("agr不能为0")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	per, err := Newperson("tom", -20)
	if err == nil {
		fmt.Printf("per: %v\n", per)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}
