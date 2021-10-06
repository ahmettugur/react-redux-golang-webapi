package business

import (
	database "onlinestorewebapi/database"
	entity "onlinestorewebapi/entities"
)

type Category struct{}

func (c Category) GetAll() (*[]entity.Category, error) {
	db, err := database.InitDB()
	if err != nil {
		return &[]entity.Category{}, err
	}

	defer database.CloseDb(db)

	var categories []entity.Category
	db.Order("Id DESC").Find(&categories)
	return &categories, nil
}

func (c Category) Get(id int) (*entity.Category, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Category{}, err
	}

	defer database.CloseDb(db)

	var category entity.Category
	db.Where("Id = ?", id).Find(&category)

	return &category, nil
}

func (c Category) Add(category *entity.Category) (*entity.Category, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Category{}, err
	}
	defer database.CloseDb(db)

	if errRecord := db.Create(&category).Error; errRecord != nil {
		return &entity.Category{}, errRecord
	}

	return category, nil
}

func (c Category) Update(category *entity.Category) (*entity.Category, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Category{}, err
	}
	defer database.CloseDb(db)

	if errRecord := db.Save(&category).Error; errRecord != nil {
		return &entity.Category{}, errRecord
	}

	/*if errRecord := db.entity(&category).Updates(entity.Category{Name: category.Name, Description: category.Description}).Error; errRecord != nil {
		return &entity.Category{}, errRecord
	}*/
	return category, nil
}

func (c Category) Delete(category entity.Category) error {
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb(db)

	if errRecord := db.Delete(&category).Error; errRecord != nil {
		return errRecord
	}

	return nil
}
