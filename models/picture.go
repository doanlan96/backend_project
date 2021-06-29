package models

type Picture struct {
	Id             uint   `json:"id"`
	ProductId      uint   `json:"product_id"`
	ProductPicture string `json:"product_picture"`
}
