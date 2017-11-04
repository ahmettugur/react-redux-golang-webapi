package entities

//ProductWithCategory struct
type ProductWithCategory struct {
	ProductId     int `gorm:"column:ProductId"`
	CategoryId    int `gorm:"column:CategoryId"`
	CategoryName  string `gorm:"column:CategoryName"`
	Name          string `gorm:"column:Name"`
	Details       string `gorm:"column:Details"`
	Price         float64 `gorm:"column:Price"`
	StockQuantity int `gorm:"column:StockQuantity"`
}
