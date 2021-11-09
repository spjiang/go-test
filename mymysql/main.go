package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
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
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

var wg sync.WaitGroup

func main() {
	wg.Add(100) // 因为有两个动作，所以增加2个计数
	for i:=0;i<100;i++ {
		go select_id()
	}
	wg.Wait() // 等待，直到计数为0
}

func select_id()  {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
	wg.Done()
}