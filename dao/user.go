package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// ExistOrNotByUserName 根据user_name查找用户是否存在
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *models.User, exist bool, err error) {
	err = dao.DB.Model(&models.User{}).Where("user_name = ?", userName).Find(&user).Error
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创造user
func (dao *UserDao) CreateUser(user *models.User) error {
	err := dao.DB.Create(&models.User{}).Create(&user).Error
	return err
}

// FindUserById 根据id获取user
func (dao *UserDao) FindUserById(id uint) (*models.User, error) {
	user := &models.User{}
	err := dao.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

// UpdateById 根据id更新user
func (dao *UserDao) UpdateById(id uint, user *models.User) error {
	return dao.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error
}
