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

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func createTable() {
	db.AutoMigrate(&User{}, &Language{})
}
func insert() {

	l := Language{
		Name: "english",
	}

	user := User{
		Languages: []Language{l},
	}
	db.Create(&user)
	db.Create(&l)
}
func main() {
	insert()
	//createTable()
}
