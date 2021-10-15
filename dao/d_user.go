package dao

import (
	"User/model"
	"gorm.io/gorm"
)

// CreateUser 新建角色
func (d *Dao)CreateUser(user *model.User) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		return err
	})
	return err
}

// SearchUser 根据keyType检索用户，keyType必须是预设类型之一
func (d *Dao)SearchUser(searchKey, keyType string) (*[]model.User, error) {
	users := &[]model.User{}
	switch keyType {
	case model.KeyTypeAccount:
		err := d.DB.Model(model.User{}).Where("account like %?%", searchKey).Find(users).Error
		return users, err
	case model.KeyTypeUserName:
		err := d.DB.Model(model.User{}).Where("username like %?%", searchKey).Find(users).Error
		return users, err
	case model.KeyTypePhone:
		err := d.DB.Model(model.User{}).Where("phone like %?%", searchKey).Find(users).Error
		return users, err
	case model.KeyTypeEmail:
		err := d.DB.Model(model.User{}).Where("email like %?%", searchKey).Find(users).Error
		return users, err
	case model.KeyTypeVoid:
		err := d.DB.Model(model.User{}).Where("account like %?% or username like %?% or phone like %?% or email like %?%", searchKey).Find(users).Error
		return users, err
	default:
		return users, model.ErrInvalidKeyType
	}
}

// EditUser 编辑用户信息
func (d *Dao)EditUser(user *model.User) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(user).Error
		return err
	})
	return err
}

// DeleteUser 根据用户id软删除用户
func (d *Dao)DeleteUser(userID uint) error {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id=?", userID).Delete(&model.User{}).Error
		return err
	})
	return err
}