package serializer

import "web_mall/models"

type Address struct {
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func BuildAddress(item *models.Address) *Address {
	return &Address{
		UserID:  item.UserID,
		Name:    item.Name,
		Phone:   item.Phone,
		Address: item.Address,
	}
}

func BuildAddressS(items []*models.Address) (res []*Address) {
	for _, item := range items {
		address := BuildAddress(item)
		res = append(res, address)
	}
	return
}
