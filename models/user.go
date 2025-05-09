package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique not null"`
	PasswordDigest string
	Email          string
	Nickname       string
	Status         string
	Avatar         string
	Money          string
	Gender         string
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// SetPassWord 对密码进行加密
func (u *User) SetPassWord(passWord string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassWord 验证密码
func (u *User) CheckPassWord(passWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(passWord))
	return err == nil
}
