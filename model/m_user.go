package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account        string      `gorm:"account"`
	UserName       string      `gorm:"username"`
	Phone          string      `gorm:"phone"`
	Email          string      `gorm:"email"`
	HashedPassword string      `gorm:"password"`
	other          interface{} `gorm:"-"`
}
