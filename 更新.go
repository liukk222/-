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
	dsn := "root:2003925liu@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
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
func updaete() {
	var user User
	db.First(&user)
	user.Name = "big tom"
	user.Age = 100
	db.Save(&user)
}
func updaete2() {
	db.Model(&User{}).Where("active=?", true).Update("name", "hello")
}
func updaete3() {
	var user User
	db.First(&user)

	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("osk")
	return nil
}

func main() {
	updaete3()
	//updaete()
	//	insert()
}
