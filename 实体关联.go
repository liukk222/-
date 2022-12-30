package main

import (
	"fmt"
	"log"

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

type Email struct {
	gorm.Model
	Email  string
	UserID int
}
type User struct {
	gorm.Model
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []Language `gorm:"many2many:user_languages;"`
}
type Address struct {
	gorm.Model
	Addresse string
	UserID   int
}
type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

func create() {
	db.AutoMigrate(&User{}, &Email{}, &Address{}, &Language{})
}
func tesd1() {
	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Addresse: "Billing Address - Address 1"},
		ShippingAddress: Address{Addresse: "Shipping Address - Address 1"},
		Emails: []Email{
			{Email: "jinzhu@example.com"},
			{Email: "jinzhu-2@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	db.Create(&user)
}
func tesd2() {
	var uesr User
	db.First(&uesr)
	var languages []Language
	db.Model(&uesr).Association("Languages").Find(&languages)
	fmt.Printf("languages: %v\n", languages)
}
func tesd3() {
	var uesr User
	db.First(&uesr)

	i := db.Model(&uesr).Association("Languages").Count()
	fmt.Printf("i: %v\n", i)
}

func main() {
	tesd3()
	//create()
}
