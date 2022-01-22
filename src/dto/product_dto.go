package dto

type ProductDto struct {
	Name string
	Id   int
}

func DummyProduct() ProductDto {
	return ProductDto{
		Name: "hello product!",
		Id:   12,
	}
}
