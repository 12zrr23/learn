package models

import (
	"fmt"
	"test/sql"
	//"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Username   string
	Password   string
	Email      string
	Phone      string
	//CreateTime time.Time
	//UpdateTime time.Time
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func CreateUser(user UserBasic) *gorm.DB {
	return sql.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return sql.DB.Delete(&user)
}

func GetUserList() []*UserBasic{
	userlist := make([]*UserBasic, 10)
	for _,v :=range userlist{
		fmt.Println(v)
	}
	return userlist
}