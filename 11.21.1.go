package main

import (
	"fmt"
	"os"
)

func write() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	f.Write([]byte(" hello golang"))
	f.Close()
}

func writeString() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteString("hello java")
	f.Close()

}

func writeAt() {
	f, err := os.OpenFile("a.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}

func main() {
	//write()
	//writeString()
	writeAt()
}

/* package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
} */
