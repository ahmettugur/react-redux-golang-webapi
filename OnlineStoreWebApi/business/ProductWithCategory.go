package business

import (
	"log"
	database "onlinestorewebapi/database"
	entity "onlinestorewebapi/entities"
)

type ProductWithCategory struct{}

func (pwc ProductWithCategory) GetAll() (*[]entity.ProductWithCategory, error) {
	db, err := database.InitDB()

	if err != nil {
		return &[]entity.ProductWithCategory{}, err
	}
	defer database.CloseDb(db)

	var productWithCategory []entity.ProductWithCategory

	if err = db.Table("Products").Select("Products.Id AS ProductId,Products.CategoryId,Categories.Name AS CategoryName,Products.Name,Products.Details,Products.Price,Products.StockQuantity").
		Joins("JOIN Categories on Categories.Id=Products.CategoryId").Order("Products.Id DESC").
		Scan(&productWithCategory).Error; err != nil {
		log.Fatal(err)
	}

	return &productWithCategory, nil

}
