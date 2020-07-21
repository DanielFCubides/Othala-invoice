package repositories

import "othala/app/products"

type ProductWriter interface {
	Create(product products.Product) (*products.Product, error)
	CreateBatch(products []products.Product) (*products.Product, error)
	Update(product products.Product) (*products.Product, error)
	Delete(product products.Product) (*products.Product, error)
}
