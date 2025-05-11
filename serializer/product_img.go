package serializer

import (
	"web_mall/config"
	"web_mall/models"
)

type ProductImage struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImage(item *models.ProductImage) ProductImage {
	return ProductImage{
		ProductId: item.ProductId,
		ImgPath:   config.Host + config.HttpPort + config.ProductPath + item.ImgPath,
	}
}

func BuildProductImages(item []*models.ProductImage) (productImgs []ProductImage) {
	for _, item := range item {
		productImg := BuildProductImage(item)
		productImgs = append(productImgs, productImg)
	}
	return
}
