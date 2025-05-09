package serializer

import (
	"web_mall/config"
	"web_mall/models"
)

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Gender   string `json:"gender"`
	Money    string `json:"money"`
	Avatar   string `json:"avatar"`
}

func BuildUser(user *models.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		Gender:   user.Gender,
		Money:    user.Money,
		Avatar:   user.Avatar + config.AvatarPath + config.Host + config.HttpPort,
	}
}
