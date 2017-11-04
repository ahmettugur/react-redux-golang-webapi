package entities

//  Usesr Struct
type User struct {
	UserId   int `gorm:"primary_key;column:UserId"`
	FullName string `gorm:"column:FullName"`
	Password string `gorm:"column:Password"`
	Email    string `gorm:"column:Email"`
	Roles []Role
}
