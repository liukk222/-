/* package main

import "fmt"

func main() {
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i:%T, %v\n", i, i)
	type MyInt2 = int
	var a MyInt2
	a = 100
	fmt.Printf("i:%T %v\n", a, a)
}
*/

/* package main

import "fmt"

func main() {
	type person struct {
		id    int
		name  string
		age   int
		email string
	}
	var tom person
	tom = person{
		id:    101,
		name:  "tom",
		age:   20,
		email: "tom@gmail.com",
	}
	fmt.Printf("tom: %v\n", tom)
}
*/

/* package main

import "fmt"

func test1() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name
	fmt.Printf("name: %v\n", name)
	fmt.Printf("p_name: %v\n", p_name)
	fmt.Printf("p_name: %v\n", *p_name)
}

func test2() {
	type person struct {
		id   int
		name string
		age  int
	}
	tom := person{
		id:   101,
		name: "tom",
		age:  20,
	}
	var p_person *person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func test3() {
	type person struct {
		id   int
		name string
		age  int
	}
	var tom = new(person)
	tom.id = 101
	tom.name = "tom"
	tom.age = 20
	fmt.Printf("tom: %v\n", *tom)
}

func main() {
	//test1()
	//test2()
	test3()
}
*/

/* package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
	per.id = 102
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	tom := Person{
		id:   101,
		name: "tom",
	}
	per := &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("---------------\n")
	showPerson2(per)
	fmt.Printf("tom: %v\n", *per)

	/* fmt.Printf("tom: %v\n", tom)
	fmt.Println("---------------\n")
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom) */
//}
