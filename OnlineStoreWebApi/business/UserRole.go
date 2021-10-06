package business

import (
	database "onlinestorewebapi/database"
	entity "onlinestorewebapi/entities"
)

type UserRole struct{}

func (ur UserRole) GetAll() (*[]entity.UserRole, error) {
	db, err := database.InitDB()
	if err != nil {
		return &[]entity.UserRole{}, err
	}

	defer database.CloseDb(db)

	var userRoles []entity.UserRole

	if errUserRoles := db.Find(&userRoles).Error; errUserRoles != nil {
		return &[]entity.UserRole{}, errUserRoles
	}

	return &userRoles, err
}

func (ur UserRole) Get(id int) (*entity.UserRole, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.UserRole{}, err
	}

	defer database.CloseDb(db)

	var userRole entity.UserRole

	if errUserRole := db.Where("Id = ?", id).Find(&userRole).Error; errUserRole != nil {
		return &entity.UserRole{}, errUserRole
	}

	return &userRole, nil
}
func (ur UserRole) GetByUserId(UserId int) (*entity.UserRole, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.UserRole{}, err
	}

	defer database.CloseDb(db)

	var userRole entity.UserRole

	if errUserRole := db.Where("UserId = ?", UserId).Find(&userRole).Error; errUserRole != nil {
		return &entity.UserRole{}, errUserRole
	}

	return &userRole, nil
}

func (ur UserRole) Add(userRole entity.UserRole) (*entity.UserRole, error) {
	db, err := database.InitDB()
	if err != nil {
		return &userRole, err
	}

	defer database.CloseDb(db)

	if errReccord := db.Create(&userRole).Error; errReccord != nil {
		return &userRole, errReccord
	}
	return &userRole, nil
}

func (ur UserRole) Update(userRole entity.UserRole) (*entity.UserRole, error) {
	db, err := database.InitDB()
	if err != nil {
		return &userRole, err
	}

	defer database.CloseDb(db)

	if errRecord := db.Update(&userRole).Error; errRecord != nil {
		return &userRole, errRecord
	}

	return &userRole, nil
}

func (ur UserRole) Delete(userRole entity.UserRole) error {
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb(db)

	if errDeleted := db.Delete(&userRole).Error; errDeleted != nil {
		return errDeleted
	}

	return nil
}
