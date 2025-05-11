package dao

import "web_mall/models"

func MigrationSeparate() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		Migrator().AutoMigrate(&models.Address{}, &models.Order{}, &models.Product{}, &models.User{},
		&models.Admin{}, &models.Carousel{}, &models.Cart{}, &models.Favorites{},
		&models.Favorite{}, &models.Notice{}, &models.ProductImage{})
	if err != nil {
		panic(err)
	}
	return
}
