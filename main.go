package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id   int64
	Name string
	Age  int
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(db:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("うまく動きました。")

}
