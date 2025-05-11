package serializer

import (
	"context"
	"web_mall/dao"
	"web_mall/models"
)

type Favorites struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DisCountPrice string `json:"discount_price"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
}

func BuildFavorite(favorite *models.Favorite, product *models.Product, user *models.User) *Favorites {
	return &Favorites{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.ProductName,
		CategoryId:    product.Category,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DisCountPrice: product.DiscountPrice,
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossId:        favorite.BossId,
	}
}

func BuildFavorites(ctx context.Context, items []*models.Favorite) (favorites []Favorites) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.ShowProductByID(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.FindUserById(item.UserId)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, *favorite)
	}
	return favorites
}
