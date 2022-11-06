package main

import "fmt"

func f1() {
	grade := 'A'
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
	default:
		fmt.Println("其他")

	}
}
func f2() {
	day := 6
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	default:
		fmt.Println("休息日")
	}
}
func f3() {
	socre := 100

	switch {
	case 80 > socre && socre >= 60:
		fmt.Println("及格")
	case socre >= 80 && socre < 90:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}
}
func main() {
	//f1()
	//f2()
	f3()
}
