package main

import (
	"fmt"
	"os"
)

func openClose() {
	/* f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	} */
	f, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}
}

func creaFile() {
	f, err := os.Create("a2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
	f2, err2 := os.CreateTemp("", "temp")
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}
	f.Close()
	f2.Close()
}

func readOps() {
	/* f, err := os.Open("tset.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for {
		buf := make([]byte, 1)
		_, err2 := f.Read(buf)
		if err2 == io.EOF {
			break
		} else {
			//fmt.Printf("n: %v\n", n)
			fmt.Printf("%v", string(buf))
		}
	}
	f.Close() */
	/* f, err := os.Open("a1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	buf := make([]byte, 4)
	_, err2 := f.ReadAt(buf, 3)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("string(buf): %v\n", string(buf))
	} */

	/* de, err := os.ReadDir("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	} */

	f, err := os.Open("a1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Seek(3, 0)
	buf := make([]byte, 10)
	_, err2 := f.Read(buf)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}

func main() {
	//openClose()
	//creaFile()
	readOps()
}
