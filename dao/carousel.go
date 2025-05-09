package dao

import (
	"context"
	"gorm.io/gorm"
	"web_mall/models"
)

type CarouselDao struct {
	*gorm.DB
}

func NewNCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// FindCarouselById 根据id获取notice
func (dao *CarouselDao) FindCarouselById() (carousel []models.Carousel, err error) {
	err = dao.DB.Model(&models.Carousel{}).Find(&carousel).Error
	return carousel, err
}
