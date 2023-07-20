package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"test.com/Config"
	"test.com/Models"
	"test.com/Routers"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
