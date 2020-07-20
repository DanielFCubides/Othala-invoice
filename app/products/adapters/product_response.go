package adapters

import "othala/app/products"

type ProductResponse struct {
	Name     string
	Category string
	Type     string
}

func (*ProductResponse) new() *ProductResponse {
	return &ProductResponse{}
}

func (p *ProductResponse) buildFromProduct(product products.Product) *ProductResponse {
	return &ProductResponse{
		Name:     product.Name,
		Category: product.Category,
		Type:     product.Type,
	}
}
