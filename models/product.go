package models

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"web_mall/cache"
)

type Product struct {
	gorm.Model
	ProductName   string
	Category      uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:true"`
	Num           int  `gorm:"default:0"`
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(context.Background(), cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加预览
func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(context.Background(), cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(context.Background(), cache.RankKey, 1, strconv.Itoa(int(product.ID)))

}
