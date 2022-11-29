package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func test1() {
	Person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	b, err := xml.MarshalIndent(Person, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func Unmarshal() {
	s := `
 <person>
 <name>tom</name>
 <age>20</age>
 <email>tom@gmail.com</email>
 </person>
`
	b := []byte(s)
	var per Person
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per)
}

func Read() {
	b, err := ioutil.ReadFile("a.xml")
	if err != nil {
		log.Fatal(err)
	}
	var per Person
	xml.Unmarshal(b, &per)
	fmt.Printf("per: %v\n", per)
}

func write() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777)
	defer f.Close()
	/* b, err := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("b: %v\n", string(b))
	if err != nil {
		log.Fatal(err)
	}
	*/
	/* s := Person{}
	err2 := xml.Unmarshal(b, &s)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} */
	//fmt.Printf("as: %v\n", as)
	e := xml.NewEncoder(f)
	e.Indent(" ", "  ")
	e.Encode(p)

}

func main() {
	write()
}

/* <person>
   <name>tom</name>
   <age>20</age>
   <email>tom@gmail.com</email>
 </person> */
