package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:2003925liu@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	id       int
	username string
	passwrod string
}

func update() {
	s := "update user_tbl set username=?,password=? where id=?"
	r, err := db.Exec(s, "big kite", "456789", 2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, err2 := r.RowsAffected()
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		} else {
			fmt.Printf("i: %v\n", i)
		}
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化错误,err:%v\n", err)
		return
	} else {
		fmt.Println("初始化成功")
	}
	update()
}
