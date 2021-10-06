package models

import entity "onlinestorewebapi/entities"

type ProductResponse struct {
	CurrentCategory int
	CurrentPage     int
	PageCount       int
	PageSize        int
	Products        []entity.Product
}
