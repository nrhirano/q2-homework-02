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
	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = engine.CreateTables(User{})
	if err != nil {
		log.Fatalf("テーブルの生成に失敗しました。: %v", err)
	}
	fmt.Println("テーブル作成が成功しました。")

	user1 := User{
		Id:   1,
		Name: "tanaka",
		Age:  20,
	}
	_, err = engine.Table("user").Insert(user1)
	if err != nil {
		log.Fatalf("テーブルの生成に失敗しました。: %v", err)
	}

	fmt.Println("うまく動きました。")

}
