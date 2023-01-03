package main

import (
	"fmt"
	"log"
	"time"

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
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

func createTable() {
	db.AutoMigrate(&User{})
}
func insert() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
		Active:   true,
	}
	db.Create(&user)
}

func del1() {
	var uesr User
	db.First(&uesr)
	db.Delete(&uesr)
}
func del2() {
	db.Delete(&User{}, 2)
}
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("delete")
	return nil
}
func main() {
	del2()
}
