package entities

//Role struct
type Role struct {
	Id   int `gorm:"primary_key;column:Id"`
	Name string `gorm:"column:Name"`
}

