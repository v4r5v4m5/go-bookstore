package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // helps the other files to interact with the db
)

func Connect() {
	d, err := gorm.Open("mysql", "gosql:sql@123/table?charset=utf8&parseTime=Trur&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
