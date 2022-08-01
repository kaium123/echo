package database

import (
	"fmt"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kaium123/practice/config"
	"github.com/kaium123/practice/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dbConfig := config.Db()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	database, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	database.Exec("USE company")

	fmt.Println("created")

	database.AutoMigrate(&model.Product{})
	db = database
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
