package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid uint `gorm:"not null"`
	User User `gorm:"foreignKey:Uid"`
	Title string `gorm:"not null"`
	Comment string `gorm:"not null"`
	Status uint8 `json:"status" gorm:"default:0;oneof=0 1 2"`
}
