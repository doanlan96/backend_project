package models

type Cart struct {
	Id        uint     `json:"id"`
	ProductId uint     `json:"product_id"`
	Quantity  uint     `json:"quantity"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductId"`
}
