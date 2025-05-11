package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

func NewFavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

// ListFavoritesByUid 查询用户个人收藏夹
func (dao *FavoritesDao) ListFavoritesByUid(uid uint) (res []*models.Favorite, err error) {
	err = dao.DB.Model(&models.Favorite{}).Where("user_id = ?", uid).Find(&res).Error
	return
}

// FavoriteExistOrNot 查找商品是否已收藏
func (dao *FavoritesDao) FavoriteExistOrNot(pid, uid uint) (exist bool, err error) {
	var cnt int64
	err = dao.DB.Model(&models.Favorite{}).Where("pid = ? AND user_id = ?", pid, uid).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	if cnt == 0 {
		return false, err
	}
	return true, nil
}

// CreateFavorites 创建收藏夹
func (dao *FavoritesDao) CreateFavorites(favorite *models.Favorite) (err error) {
	err = dao.DB.Model(&models.Favorite{}).Create(favorite).Error
	return
}

// DeleteFavorites 删除收藏夹
func (dao *FavoritesDao) DeleteFavorites(fid uint, uid uint) (err error) {
	return dao.DB.Model(&models.Favorite{}).Where("user_id = ? AND id = ?", uid, fid).Delete(&models.Favorite{}).Error
}
