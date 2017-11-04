package business

import (
	database "../database"
	entity "../entities"
	"fmt"
	"log"
)

type User struct {}

func (u User)GetAll()(*[]entity.User,error)  {
	db,err:=database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return &[]entity.User{},err
	}

	defer database.CloseDb(db)

	var users []entity.User

	if errUsers:=db.Find(&users).Error;errUsers!= nil {
		return &users,err
	}


	return &users,nil;
}

func (u User)GetUserRoles(userId int)(*[]entity.Role,error)  {
	db,err:=database.InitDB()
	if err != nil {
		return &[]entity.Role{},err
	}

	defer database.CloseDb(db)

	var userRoles []entity.Role

	if err = db.Joins("JOIN UserRoles on UserRoles.RoleId=Roles.Id").
		Joins("JOIN Users on UserRoles.UserId=Users.UserId").Where("Users.UserId=?", userId).
		Group("Roles.Name,Roles.Id").Find(&userRoles).Error; err != nil {
		log.Fatal(err)
	}


	return &userRoles,nil;
}

func (u User) Get(id int) (*entity.User,error) {
	db,err:=database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return &entity.User{},err
	}

	defer  database.CloseDb(db)

	var user entity.User


	if errUser:=db.Where("UserId = ?",id).Find(&user).Error;errUser!= nil {
		return &user,err
	}

	return &user,nil;
}

func (u User) ValidateUser(email string,password string) (*entity.User,error) {
	db,err:=database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return nil,err
	}

	defer  database.CloseDb(db)

	var user entity.User


	if errUser:=db.Where("Email = ? and Password = ? ",email,password).Find(&user).Error;errUser!= nil {
		return nil,err
	}

	return &user,nil;
}

func(u User) Add(user entity.User) (*entity.User,error)  {

	db,err:=database.InitDB()
	if err != nil {
		return &entity.User{},err
	}

	defer database.CloseDb(db)

	if errRecord:=db.Create(&user).Error; errRecord!=nil {
		return &entity.User{},errRecord
	}

	return &user,nil
}

func (u User)Update (user entity.User) (*entity.User,error)  {
	db,err:=database.InitDB()
	if err !=nil{
		return &entity.User{},err
	}

	defer database.CloseDb(db)

	/*if errRecord:=db.entity(&user).Updates(entity.User{FullName:user.FullName,Email:user.Email,Password:user.Password}).Error;errRecord != nil {
		return entity.User{},errRecord
	}*/
	if errRecord := db.Save(&user).Error; errRecord != nil {
		return &entity.User{}, errRecord
	}

	return &user,nil
}

func (u User) Delete(user entity.User) error  {
	db,err:=database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb(db)

	if errDeleted:=db.Delete(&user).Error;errDeleted!=nil{
		return errDeleted
	}

	return nil
}