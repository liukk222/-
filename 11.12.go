/* package main

import "fmt"

type person struct {
	name string
}

func (per person) eat() {
	fmt.Printf("%v,eat...", per.name)
}

func (per person) sleep() {
	fmt.Printf("%v,sleep...", per.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	cus := Customer{
		name: "tom",
	}
	b := cus.login("tom", "123")
	fmt.Printf("b: %v\n", b)

	/* per := person{
		name: "tom",
	}
	per.eat()
	per.sleep() */
//}

package main

import "fmt"

func tset1() {
	type person struct {
		name string
	}
	p1 := person{
		name: "tom",
	}
	p2 := &person{
		name: "tom",
	}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
}

type Person struct {
	name string
}

func showPerson1(per Person) {
	per.name = "tom..."
}

func showPerson2(per *Person) {
	per.name = "tom..."
}

func (per Person) showPerson3() {
	per.name = "tom..."
}

func (per *Person) showPerson4() {
	per.name = "tom..."
}

func main() {
	/* type person struct {
		name string
	}
	p1 := person{
		name: "tom",
	}
	p2 := &person{
		name: "tom",
	}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2) */

	p1 := Person{
		name: "tom",
	}
	p2 := &Person{
		name: "tom",
	}
	/*
		showPerson1(p1)
		fmt.Printf("p1: %v\n", p1)
		showPerson2(p2)
		fmt.Printf("p2: %v\n", *p2) */
	p1.showPerson3()
	p2.showPerson4()

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", *p2)
}
