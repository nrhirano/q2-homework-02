package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	ID   int64  `xorm:"id pk autoincr"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}

func Insert(engine xorm.Engine) {
	user := User{
		ID:   1,
		Name: "tanaka",
		Age:  20,
	}

	_, err := engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("< Insert > ID:%d 名前:%s, 年齢:%d\n", user.ID, user.Name, user.Age)
}

func Get(engine xorm.Engine) {
	user := User{}
	result, err := engine.Where("ID = ?", 1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("Not Found")
	}
	fmt.Printf("< Get > ID:%d 名前:%s, 年齢:%d\n", user.ID, user.Name, user.Age)
}

func CreateTable(engine xorm.Engine) {
	err := engine.CreateTables(User{})
	if err != nil {
		log.Fatalf("テーブルの生成に失敗しました。: %v", err)
	}
	fmt.Println("テーブル作成が成功しました。")
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(db:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World!")

	CreateTable(*engine)
	Insert(*engine)
	Get(*engine)

	defer engine.Close()
}
