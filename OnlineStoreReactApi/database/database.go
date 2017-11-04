package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func InitDB() (*gorm.DB, error) {
	connString := "sqlserver://sa:admin@localhost:52616?database=SimpleStore&connection+timeout=30"
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
