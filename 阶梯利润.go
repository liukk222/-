package main

import "fmt"

/* 企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，
低于 10 万元的部分按 10% 提成，
高于 10 万元的部分，可提成 7.5%。
20 万到 40 万之间时，高于 20 万元的部分，
可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，
高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。

从键盘输入当月利润 I，求应发放奖金总数？ */

func main() {
	var i float32 = 0.0
	var b float32 = 0.0
	fmt.Println("请输入利润为XX万元,不用输入后面万数位")
	fmt.Scanf("%f\n", &i)
	switch {
	case i > 100:
		b = (i - 100) * 0.01
		i = 100
		fallthrough
	case i > 60:
		b = b + (i-60)*0.015
		i = 60
		fallthrough
	case i > 40:
		b = b + (i-40)*0.03
		i = 40
		fallthrough
	case i > 20:
		b = b + (i-20)*0.05
		i = 20
		fallthrough
	case i > 10:
		b = b + (i-10)*0.075
		i = 10
		fallthrough
	default:
		b = b + i*0.1
		break

	}

	fmt.Printf("提成: %f\n", b)

}
