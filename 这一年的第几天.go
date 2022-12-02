package main

import "fmt"

func fmterr() {
	fmt.Println("日期输入错误")
}

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Println("请输入日期：")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if ((a%4 == 0) && (a%100 != 0)) || (a%400 == 0) {
		switch {
		case b == 1:
			if c <= 31 {
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 2:
			if c <= 29 {
				c = c + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 3:
			if c <= 31 {
				c = c + 29 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 4:
			if c <= 30 {
				c = c + 29 + 31 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 5:
			if c <= 31 {
				c = c + 29 + 31 + 31 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 6:
			if c <= 30 {
				c = c + 29 + 31 + 31 + 30 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 7:
			if c <= 31 {
				c = c + 29 + 31 + 31 + 30 + 31 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 8:
			if c <= 31 {
				c = c + 213
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 9:
			if c <= 30 {
				c = c + 213 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 10:
			if c <= 31 {
				c = c + 213 + 61
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 11:
			if c <= 30 {
				c = c + 213 + 92
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 12:
			if c <= 31 {
				c = c + 213 + 92 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		}

	} else {
		switch {
		case b == 1:
			if c <= 31 {
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 2:
			if c <= 28 {
				c = c + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 3:
			if c <= 31 {
				c = c + 28 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 4:
			if c <= 30 {
				c = c + 28 + 31 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 5:
			if c <= 31 {
				c = c + 28 + 31 + 31 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 6:
			if c <= 30 {
				c = c + 28 + 31 + 31 + 30 + 31
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 7:
			if c <= 31 {
				c = c + 28 + 31 + 31 + 30 + 31 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 8:
			if c <= 31 {
				c = c + 212
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 9:
			if c <= 30 {
				c = c + 213 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 10:
			if c <= 31 {
				c = c + 213 + 60
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 11:
			if c <= 30 {
				c = c + 213 + 91
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		case b == 12:
			if c <= 31 {
				c = c + 213 + 91 + 30
				fmt.Printf("这是一年的第%d天\n", c)
			} else {
				fmterr()
			}
		}
	}

}
