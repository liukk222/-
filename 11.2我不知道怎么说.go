/* package main */
/*  */
/* import ( */
/* 	"fmt" */
/* 	"math" */
/* 	"unsafe" */
/* ) */
/*  */
/* func main() { */
/*  */

/* package main */
/*  */
/* import ( */
/* 	"fmt" */
/* 	"math" */
/* ) */
/*  */
/* func main() { */
/* 	fmt.Printf("%f\n", math.Pi) */
/* 	fmt.Printf("%.2f\n", math.Pi) */
/* } */

/* // 十进制 */
/* var a int = 10 */
/* fmt.Printf("%d \n", a) // 10 */
/* fmt.Printf("%b \n", a) // 10 */ //10  //占位符%b表示二进制

/* // 十六进制  以0x开头 */
/* var c int = 0xff */
/* fmt.Printf("%x \n", c) // ff */
/* fmt.Printf("%X \n", c) // FF */
//}

/* var i8 int8 */
/* var i16 int16 */
/* var i32 int32 */
/* var i64 int64 */
/* var ui8 uint8 */
/* var ui16 uint16 */
/* var ui32 uint32 */
/* var ui64 uint64 */

/* fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8) */
/* fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16) */
/* fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32) */
/* fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64) */

/* fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8) */
/* fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16) */
/* fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32) */
/* fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64)) */

/* var f32 float32 */
/* var f64 float64 */

/* fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32) */
/* fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64) */

/* var ui uint */
/* ui = uint(math.MaxUint64) //再+1会导致overflows错误 */
/* fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui) */

/* var imax, imin int */
/* imax = int(math.MaxInt64) //再+1会导致overflows错误 */
/* imin = int(math.MinInt64) //再-1会导致overflows错误 */

/* fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax) */
//}
package main

import (
	"fmt"
	"strings"
)

/* (
	"fmt"
	"strings"
) */

func main() {
	//转义字符
	s := "hello,world"
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, ","))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "a"))
	/* a := 2
	b := 5
	fmt.Printf("s: %v\n", s[a:])
	fmt.Printf("s: %v\n", s[a:b])
	fmt.Printf("s[:b]: %v\n", s[:b])
	fmt.Printf("len(s): %v\n", len(s)) */
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))
	/* var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String()) */
	/* name := "tom"
	age := "20"
	s := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", s) */

	/* s1 := "tom"
	s2 := "20"
	msg := fmt.Sprintf("%s,%s", s1, s2)
	fmt.Printf("msg: %v\n", msg) */

	/* var s string = "hello word" */
	/* fmt.Printf("s: %v\n", s) */
	/* s3 := ` */
	/* a1 */
	/* a2 */
	/* a3 */
	/* ` */
	/* fmt.Printf("s3: %v\n", s3) */
}
