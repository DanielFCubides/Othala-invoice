package repositories

import "othala/app/products"

type ProductReader interface {
	FindById(id string) (*products.Product, error)
	FindAll(searchParams []string) ([]*products.Product, error)
}
