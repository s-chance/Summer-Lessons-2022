package db

import (
	"fmt"
	"log"

	"demo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.UserName,
		config.Config.DB.Password,
		config.Config.DB.Address,
		config.Config.DB.DBName,
	)
	var err error
	if config.Config.Dev {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		DB, err = gorm.Open(mysql.Open(dsn))
	}
	if err != nil {
		log.Panicln("Database Error: ", err)
	}
}
