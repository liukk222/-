package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("hello world,我们")
	f, err2 := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	if err2 != nil {
		log.Fatal(err2)
	}
	_, err := io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func createFile() {
	// 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 第一个参数 目录默认：Temp 第二个参数 文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

func main() {
	//testCopy()
	/* r := strings.NewReader("hello,word!")

	for {
		b := make([]byte, 3)
		 _, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%v", string(b))

	} */
	createFile()
}

/* package main

import (
	"fmt"
	"os"
)

func main() {
	// s := os.Getenv("GOPATH")
	// // 	fmt.Printf("s: %v\n", s) */
// s, b := os.LookupEnv("GOROOT")
// if b {
// 	fmt.Printf("s: %v\n", s)
// } else {
// 	fmt.Printf("b: %v\n", b)
//}
/*s2 := os.Getenv("GOPATH")
	fmt.Printf("s2: %v\n", s2)
}
*/
