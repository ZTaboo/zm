package db

import (
	"ZM/db/dbModel"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var Db *gorm.DB

func init() {
	//	 获取user目录
	userDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	//	 在user目录下的zm.db文件
	dbFile := userDir + "/zm.db"
	if db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		// 日志级别 silent
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		log.Println(err)
		return
	} else {
		Db = db
	}
	err = Db.AutoMigrate(
		&dbModel.User{},
		&dbModel.Task{},
	)
	if err != nil {
		log.Println(err)
		return
	}
}
