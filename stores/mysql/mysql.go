package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func NewDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:123456@/my_note?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql init error,msg:%v", err)
		return nil, err
	}

	return db, nil
}
