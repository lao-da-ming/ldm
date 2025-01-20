package Initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitMysql() *gorm.DB {
	dsn := "root:123456@tcp(192.168.1.100:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接数据库失败：", err)
	}
	mdb, err := db.DB()
	if err != nil {
		log.Fatalln("连接数据库失败：", err)
	}
	mdb.SetMaxOpenConns(10)
	mdb.SetMaxIdleConns(1)
	mdb.SetConnMaxIdleTime(time.Second * 60)
	mdb.SetConnMaxLifetime(time.Minute * 3)
	fmt.Println("连接数据库成功")
	return db
}
