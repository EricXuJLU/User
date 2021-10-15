package model

import "errors"

const (
	KeyTypeAccount  = "Account"
	KeyTypeUserName = "UserName"
	KeyTypePhone    = "Phone"
	KeyTypeEmail    = "Email"
	KeyTypeVoid     = ""
)

var (
	ErrInvalidKeyType = errors.New("invalid key type")
)
