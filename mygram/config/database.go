package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/mygramdb?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
