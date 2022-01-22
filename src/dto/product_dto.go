package dto

type ProductDto struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Id     int     `json:"id"`
	ImgUrl string  `json:"imgUrl"`
}
