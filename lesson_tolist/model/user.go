package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	passCost = 12
)
type User struct {
	gorm.Model
	NickName string	`gorm:"not null;unique;varchar(65)"`
	Passwd	 string `gorm:"not null"`
	Email    string  `gorm:"null"`
}

func (u *User) CryptPasswd(passwd string) error{
	cryptedPasswd,err:=bcrypt.GenerateFromPassword([]byte(passwd),passCost)
	if err != nil {
		return err
	}
	u.Passwd=string(cryptedPasswd)
	return err
}

func (u *User) CheckPasswd(passwd string) bool {
	if err:=bcrypt.CompareHashAndPassword([]byte(u.Passwd),[]byte(passwd));err!=nil{
		return false
	}
	return true
}