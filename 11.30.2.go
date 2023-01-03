package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
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

func inser() {
	s := "insert into user_tbl(username,password)values(?,?)"
	r, err := db.Exec(s, "zhangsan", "zs123")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, err2 := r.LastInsertId()
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
		} else {
			fmt.Printf("i: %v\n", i)
		}
	}
}

func inser2(username string, passwrod string) {
	s := "insert into user_tbl(username,password)values(?,?)"
	r, err := db.Exec(s, username, passwrod)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, err2 := r.LastInsertId()
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
	//inser()
	inser2("lisi", "lisi123")
}
