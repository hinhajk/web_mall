package serializer

import "web_mall/models"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item *models.Favorites) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildListCategories(items []models.Favorites) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(&item)
		categories = append(categories, category)
	}
	return categories
}
