package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"default:''"` //名字
	Age  int32  `json:"age" gorm:"default:0"`   //年龄
}
type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}
func (u *UserModel) FindOne(ctx context.Context, id int64, rsp *User) error {
	return u.db.Model(&User{}).Where("id = ?", id).First(&rsp).Error
}
