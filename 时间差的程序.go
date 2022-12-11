package main

import (
	"fmt"
	"time"
)

func main() {
	// 设置时区为中国时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从云端读取存储的时间
	// 假设这里读取到的时间是 "2022-12-11 10:00:00"
	storedTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-12-11 10:00:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取当前时间
	now := time.Now().In(loc)

	// 计算两个时间之间的时间差
	duration := now.Sub(storedTime)

	// 判断时间差，并输出结果
	if duration < time.Hour {
		// 小于1小时，输出距离多少分钟
		fmt.Printf("距离 %d 分钟", int(duration.Minutes()))
	} else if duration < 24*time.Hour {
		// 小于24小时，输出距离多少小时
		fmt.Printf("距离 %d 小时", int(duration.Hours()))
	} else if duration < 3*24*time.Hour {
		// 小于3天，输出距离多少天
		fmt.Printf("距离 %d 天", int(duration.Hours()/24))
	} else {
		// 大于3天，输出存储的时间
		fmt.Println(storedTime.Format("2006-01-02 15:04:05"))
	}
}
