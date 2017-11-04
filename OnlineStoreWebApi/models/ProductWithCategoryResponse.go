package models

import entity "../entities"

type ProductWithCategoryResponse struct {
	PageCount int
	PageSize int
	Products []entity.ProductWithCategory
}