package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
	TestId   int    `db:"test_id"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3366)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {
	for i:=1;i<10;i++ {
		r, err := Db.Exec("insert into person(username, sex, email,test_id)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}

		fmt.Println("insert succ:", id)
	}

}