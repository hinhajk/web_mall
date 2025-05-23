package serializer

import (
	"time"
	"web_mall/config"
	"web_mall/models"
)

type ProductService struct {
	ID            uint      `json:"id"`
	ProductName   string    `json:"product_name"`
	Category      uint      `json:"category"`
	Title         string    `json:"title"`
	Info          string    `json:"info"`
	ImgPath       string    `json:"img_path"`
	Price         string    `json:"price"`
	OnSale        bool      `json:"onSale"`
	DisCountPrice string    `json:"discount_price"`
	View          uint64    `json:"view"`
	CreateTime    time.Time `json:"create_time"`
	Num           int       `json:"num"`
	BossId        uint      `json:"boss_id"`
	BossName      string    `json:"boss_name"`
	BossAvatar    string    `json:"boss_avatar"`
}

func BuildProduct(product *models.Product) *ProductService {
	return &ProductService{
		ID:            product.ID,
		ProductName:   product.ProductName,
		Category:      product.Category,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
		Price:         product.Price,
		OnSale:        product.OnSale,
		DisCountPrice: product.DiscountPrice,
		View:          product.View(),
		CreateTime:    product.CreatedAt,
		Num:           product.Num,
		BossId:        product.BossId,
		BossName:      product.BossName,
		BossAvatar:    config.Host + config.HttpPort + config.ProductPath + product.BossAvatar,
	}
}

func BuildProducts(products []*models.Product) (result []*ProductService) {
	for _, product := range products {
		result = append(result, BuildProduct(product))
	}
	return result
}
