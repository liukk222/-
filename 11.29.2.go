package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test1() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tomgmail.com",
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func test2() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tomgmail.com"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p)
}

func test3() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tomgmail.com", "Parents":["tom", "kite"]}`)
	var f map[string]interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("p: %v\n", f)
	for k, v := range f {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}

func test4() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test6() {
	f, err := os.Open("a.json")
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	d := json.NewDecoder(f)
	var v map[string]interface{}
	d.Decode(&v)
	fmt.Printf("v: %v\n", v)
}

func test7() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}
	f, err := os.OpenFile("a.json", os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	d := json.NewEncoder(f)
	d.Encode(p)

}

func main() {
	//test2()
	test7()

}
