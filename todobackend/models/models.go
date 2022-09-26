package models

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "zxcvbnm@1"
const DB_NAME = "todo"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

type Todo struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}
type UpdateTodoInput struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func InitDb() {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return
	}

	dasb, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	dasb.Close();
	Db.AutoMigrate(&Todo{})
}
