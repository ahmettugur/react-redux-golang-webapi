package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func InitDB() (*gorm.DB, error) {
	connString := "sqlserver://sa:Ahmet1990*@localhost:1453?database=Store&connection+timeout=30"
	//var err error
	db, err := gorm.Open("mssql", connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func CloseDb(db *gorm.DB) {
	db.Close()
}
