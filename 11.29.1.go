package main

import (
	"fmt"
	"time"
)

func test1() {
	t := time.Now()
	fmt.Printf("t: %T\n", t)
	fmt.Printf("t: %v\n", t)
}

func test2() {
	now := time.Now()
	fmt.Printf("now.Unix(): %T\n", now.Unix())
	fmt.Printf("now.Unix(): %v\n", now.Unix())
}

func test3() {
	t := time.Now()
	fmt.Printf("t.Format(\"2006/01/02 15:04\"): %v\n", t.Format("2006/01/02 15:04"))
}
func test4() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	//test1()
	//test2()
	test4()
}
