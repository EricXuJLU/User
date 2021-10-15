package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account        string
	UserName       string
	Phone          string
	Email          string
	HashedPassword string
	other          interface{} `gorm:"-"`
}
