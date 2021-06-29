package models

type Product struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	IsSale      bool       `json:"is_sale"`
	IsHot       bool       `json:"is_hot"`
	CategoryId  uint       `json:"category_id"`
	Picture     []*Picture `json:"pictures" gorm:"foreignKey:ProductId"`
	Price       float64    `json:"price"`
	SalePrice   float64    `json:"sale_price"`
	Category    *Category  `json:"category" gorm:"foreignKey:CategoryId"`
}
