package business

import (
	database "onlinestorewebapi/database"
	entity "onlinestorewebapi/entities"
)

type Role struct{}

func (r Role) GetAll() (*[]entity.Role, error) {
	db, err := database.InitDB()
	if err != nil {
		return &[]entity.Role{}, err
	}

	defer database.CloseDb(db)

	var roles []entity.Role

	if errRoles := db.Find(&roles).Error; errRoles != nil {
		return &roles, err
	}

	return &roles, err
}

func (r Role) Get(id int) (*entity.Role, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Role{}, nil
	}

	defer database.CloseDb(db)

	var role entity.Role

	if errRole := db.Where("Id = ?", id).Find(&role); errRole != nil {
		return &role, err
	}

	return &role, nil
}

func (r Role) Add(role entity.Role) (*entity.Role, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Role{}, nil
	}

	defer database.CloseDb(db)

	if errRecor := db.Create(&role).Error; errRecor != nil {
		return &entity.Role{}, err
	}

	return &role, nil
}

func (r Role) Update(role entity.Role) (*entity.Role, error) {
	db, err := database.InitDB()
	if err != nil {
		return &entity.Role{}, err
	}

	defer database.CloseDb(db)

	if errRecord := db.Save(&role).Error; errRecord != nil {
		return &entity.Role{}, err
	}

	return &role, nil

}

func (r Role) Delete(role entity.Role) error {
	db, err := database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb(db)

	if errDelete := db.Delete(&role).Error; errDelete != nil {
		return errDelete
	}

	return nil
}
