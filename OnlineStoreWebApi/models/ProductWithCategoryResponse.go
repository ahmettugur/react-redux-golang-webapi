package models

import entity "onlinestorewebapi/entities"

type ProductWithCategoryResponse struct {
	PageCount int
	PageSize  int
	Products  []entity.ProductWithCategory
}
