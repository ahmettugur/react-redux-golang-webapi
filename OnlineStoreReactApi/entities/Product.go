package entities

// Product struct
type Product struct {
	Id           int `gorm:"primary_key;column:Id" `
	CategoryId    int `gorm:"column:CategoryId"`
	Name          string `gorm:"column:Name"`
	Details       string `gorm:"column:Details"`
	Price         float64 `gorm:"column:Price" json:"Price,string,omitempty"`
	StockQuantity int `gorm:"column:StockQuantity" json:"StockQuantity,string,omitempty"`
}
