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
}

func createTable() {
	db.AutoMigrate(&User{})
}
func Select1() {
	var user User
	//db.First(&user)
	/* fmt.Printf("user.ID: %v\n", user.ID)
	db.Take(&user)
	fmt.Printf("user.ID: %v\n", user.ID) */
	//fmt.Printf("user: %v\n", user)
	db.Last(&user)
	fmt.Printf("user.ID: %v\n", user.ID)
}
func Select2() {
	//var user User
	/*db.First(&user, 10)//"10"
	fmt.Printf("user.ID: %v\n", user.ID) */
	var users []User
	db.Find(&users, []int{1, 5, 6})
	for _, user := range users {
		fmt.Printf("user.ID: %v\n", user.ID)
	}

}

func Select3() {
	var users []User
	resulit := db.Find(&users)
	fmt.Printf("resulit.RowsAffected: %v\n", resulit.RowsAffected)
}

func main() {

	Select3()
}
