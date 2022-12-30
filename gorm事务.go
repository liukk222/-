package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:2003925liu@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	//禁用全局事务
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ /* SkipDefaultTransaction:true */ })
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func test1() {
	//session级别的禁用事务
	db.Session(&gorm.Session{SkipDefaultTransaction: true})
}

type User struct {
	gorm.Model
	Name string
}

// 测试事务控制 没有事务控制
func test2() {
	user := User{
		Name: "tom",
	}
	db.Create(&user)
	db.Create(nil)
}

// 测试事务控制 手动控制
func test3() {
	user := User{
		Name: "tom",
	}
	d := db.Begin()

	d.Create(&user)
	err := d.Create(nil).Error
	if err != nil {
		d.Rollback()
	} else {
		d.Commit()
	}
}
func test4() {
	user := User{
		Name: "tom",
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		// if err2 := tx.Create(nil).Error; err2 != nil {
		// 	return err2
		// }
		return nil
	})
}
func CreateAnimals(db *gorm.DB) error {
	// 再唠叨一下，事务一旦开始，你就应该使用 tx 处理数据
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&User{Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func main() {
	test4()
}
