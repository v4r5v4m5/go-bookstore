package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB // helps the other files to interact with the db
)

func Connect() {
	d, err := gorm.Open("mysql", "<uname>:<pwd>@tcp(<server>:<port>)/<table>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
