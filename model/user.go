package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string //`gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"not null"`
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}
