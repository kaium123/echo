package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kaium123/practice/model"
)

var db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS company")
	db.Exec("USE company")

	fmt.Println("created")

	db.AutoMigrate(&model.Product{})

}

func GetDB() *gorm.DB {
	InitDB()
	return db
}
