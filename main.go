package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id   int64  `xorm:"id pk"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}

func Insert(engine xorm.Engine) {
	user := User{
		Id:   1,
		Name: "tanaka",
		Age:  20,
	}

	_, err := engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("< Insert > ID:%d 名前:%s, 年齢:%d\n", user.Id, user.Name, user.Age)
}

func Get(engine xorm.Engine) {
	user := User{}
	result, err := engine.Where("Id = ?", 1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("Not Found")
	}
	fmt.Printf("< Get > ID:%d 名前:%s, 年齢:%d\n", user.Id, user.Name, user.Age)
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World!")

	Insert(*engine)
	Get(*engine)

	defer engine.Close()
}
