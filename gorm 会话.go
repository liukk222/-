package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:2003925liu@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DryRun: true}) //全局配置
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
func tesd1() {
	//会话
	db.Session(&gorm.Session{DryRun: true})
}
