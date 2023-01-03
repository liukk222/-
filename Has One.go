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
func tesd1() {
	// User 有一张 CreditCard，UserID 是外键
	type CreditCard struct {
		gorm.Model
		Number string
		UserID uint
	}
	type User struct {
		gorm.Model
		CreditCard CreditCard
	}

}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

func test2() {

	db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})

}
func tesd3() {
	db.Create(&Cat{Name: "Cat1", Toy: Toy{Name: "toy2"}})
}

func main() {
	tesd3()
}
*/