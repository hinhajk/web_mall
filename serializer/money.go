package serializer

import (
	"web_mall/models"
	"web_mall/pkg/utils"
)

type Money struct {
	UserId    uint   `json:"user_id"`
	Username  string `json:"username"`
	UserMoney string `json:"user_money"`
}

func BuildMoney(user *models.User, key string) Money {
	utils.Encrypt.SetKey(key)
	return Money{
		UserId:    user.ID,
		Username:  user.UserName,
		UserMoney: utils.Encrypt.AesDecoding(user.Money),
	}
}
