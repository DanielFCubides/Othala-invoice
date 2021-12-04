package repositories

import "othala/app/products"

type Product struct {
	Name     string `gorm:"primaryKey"`
	Category string
	Type     string
}

func Migrate(){

}

func (p Product) ToDomain() products.Product {
	return products.Product{
		Name:     p.Name,
		Category: p.Category,
		Type:     p.Type,
	}
}

func (p Product) FromDomain(product products.Product) Product {
	return Product{
		Name:     product.Name,
		Category: product.Category,
		Type:     product.Type,
	}
}

