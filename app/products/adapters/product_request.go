package adapters

import "othala/app/products"

type ProductRequest struct {
	Name     string
	Category string
	Type     string
}

func (*ProductRequest) new() *ProductRequest {
	return &ProductRequest{}
}

func (request *ProductRequest) mapToProduct() products.Product {
	return products.Product{
		Name:     request.Name,
		Category: request.Category,
		Type:     request.Type,
	}
}
