package serializer

import "web_mall/models"

type Cart struct {
	ID            uint   `json:"id" form:"id"`
	UserId        uint   `json:"user_id" form:"user_id"`
	ProductId     uint   `json:"product_id" form:"product_id"`
	Name          string `json:"name" form:"name"`
	CreatedAt     int64  `json:"created_at" form:"created_at"`
	MaxNum        uint   `json:"max_num" form:"max_num"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Check         bool   `json:"check" form:"check"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	BossId        uint   `json:"boss_id" form:"boss_id"`
	BossName      string `json:"boss_name" form:"boss_name"`
	Num           uint   `json:"num" form:"num"`
}

func BuildCart(cart *models.Cart, product *models.Product, boss *models.User) *Cart {
	return &Cart{
		ID:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		Name:          product.ProductName,
		CreatedAt:     cart.CreatedAt.Unix(),
		MaxNum:        uint(product.Num),
		ImgPath:       product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: product.DiscountPrice,
		BossId:        cart.BossId,
		BossName:      boss.UserName,
		Num:           cart.Num,
	}
}

//func BuildCarts(carts []*models.Cart) (res []*Cart) {
//	for _, item := range carts {
//		res = append(res, BuildCart(item))
//	}
//	return
//}
