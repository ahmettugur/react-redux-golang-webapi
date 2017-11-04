package entities

/*import (
	"fmt"

	database "../database"
)*/

//Category struct
type Category struct {
	Id          int    `gorm:"primary_key;column:Id"`
	Name        string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
}
/*
func (c Category) GetAll() []Category {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
	}

	defer database.CloseDb(db)

	var categories []Category
	db.HasTable(&categories)
	db.Find(&categories)
	fmt.Println(categories)
	return categories
}

func (c Category) Get(id int) Category {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return Category{}
	}

	defer database.CloseDb(db)

	var category Category
	db.Where("Id = ?", id).Find(&category)
	fmt.Println(category)

	return category
}

func (c Category) Add(category Category) (Category, error) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return Category{}, err
	}
	defer database.CloseDb(db)

	if errRecord := db.Create(&category).Error; errRecord != nil {
		return Category{}, errRecord
	}

	return category, nil
}

func (c Category) Update(category Category) (Category, error) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
		return Category{}, err
	}
	defer database.CloseDb(db)

	if errRecord := db.Model(&category).Updates(Category{Name: category.Name, Description: category.Description}).Error; errRecord != nil {
		return Category{}, errRecord
	}
	return category, nil
}

func (c Category) Delete(category Category) error {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("No connect database")
	}

	defer database.CloseDb(db)

	if errRecord := db.Delete(&category).Error; errRecord != nil {
		return errRecord
	}

	return nil
}
*/