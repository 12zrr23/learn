package sql

import (
	"log"
	"time"
	"os"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB

func InitSql(){
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB,_ =	gorm.Open(mysql.Open("root:my-secret-pw@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"),
	&gorm.Config{Logger:newlogger})
	fmt.Println(" MySQL inited 。。。。")
}