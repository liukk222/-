package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

type User struct {
	gorm.Model
	CreditCards []CreditCards
}
type CreditCards struct {
	gorm.Model
	Number string
	UserID uint
}

func createTable() {
	db.AutoMigrate(&User{}, &CreditCards{})

}

func main() {
	createTable()
}
