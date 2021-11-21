package usecases

import "othala/app/products"

type ProductManager interface {
	GetAll() ([]products.Product, error)
	GetById(productId string) (*products.Product, error)
	Delete(productId string) (*products.Product, error)
	Create(product products.Product) (*products.Product, error)
	Update(product products.Product) (*products.Product, error)
}
