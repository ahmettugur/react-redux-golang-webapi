package business

import (
	database "../database"
	entity "../entities"
	"fmt"
)

type Product struct {}

func (p Product) GatAll(categoryId int,page int ) (*[]entity.Product,error)  {
	db,err:=database.InitDB()
	if err != nil{
		return &[]entity.Product{},err
	}

	defer database.CloseDb(db)

	var products []entity.Product
	if categoryId != 0 {
		db.Where("CategoryId=? ",categoryId).Order("Id DESC").Find(&products)
	}else{
		db.Order("Id DESC").Find(&products)
	}

	return &products,nil
}

func (p Product) Get(id int) (*entity.Product,error) {
	db,err:=database.InitDB()
	if err != nil{
		return &entity.Product{},err
	}

	var product entity.Product

	db.Where("Id = ?",id).Find(&product)
	fmt.Println(product)
	return &product,nil
}

func (p Product)Add(product *entity.Product) (*entity.Product,error) {
	db,err:=database.InitDB()
	if err != nil{
		 return &entity.Product{},err
	}

	defer  database.CloseDb(db);
	if errRecord := db.Create(&product).Error; errRecord != nil {
		return &entity.Product{},errRecord
	}

	return product,nil
}

func (p Product) Update(product *entity.Product) (*entity.Product,error)  {
	db,err:=database.InitDB()
	if err != nil {
		return &entity.Product{},err
	}

	defer database.CloseDb(db)

	if errRecord:=db.Save(&product).Error;errRecord !=nil {
		return &entity.Product{},err
	}

	return product,nil
}

func (p Product) Delete(product entity.Product) error  {
	db,err:=database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb(db)

	if errDeleted:=db.Delete(&product).Error;errDeleted!=nil{
		return errDeleted
	}

	return nil
}