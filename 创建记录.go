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
}

func createTable() {
	db.AutoMigrate(&User{})
}

func intsert1() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
	}
	resulit := db.Create(&user)
	fmt.Printf("resulit.RowsAffected: %v\n", resulit.RowsAffected)
	fmt.Printf("user.ID: %v\n", user.ID)
}
func intsert2() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
	}
	db.Select("Name", "Age", "CreatedAt").Create(&user)
}

func intsert3() {
	users := []User{{Name: "jinzhu1", Birthday: time.Now()}, {Name: "jinzhu2", Birthday: time.Now()}, {Name: "jinzhu3", Birthday: time.Now()}}
	db.Create(&users)

	for _, d := range users {
		fmt.Printf("d.ID: %v\n", d.ID) // 1,2,3
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return
}
func intsert4() {
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "as", "Age": 18,
	})
}

func main() {
	//createTable()
	intsert4()
}
