package implementations

import (
	"othala/app/config"
	"othala/app/products"
	"othala/app/products/usecases"
)

type ProductManagerImpl struct{}

func init() {
	if err := config.Injector.Provide(newProductManagerImpl); err != nil {
		panic("error providing product manager")
	}
}

func newProductManagerImpl() usecases.ProductManager {
	return &ProductManagerImpl{}
}

func (manager *ProductManagerImpl) GetAll() ([]products.Product, error) {
	return []products.Product{{
		Name:     "club colombia roja",
		Category: "alcohol",
		Type:     "beer",
	}}, nil
}

func (manager *ProductManagerImpl) GetById(productId string) (products.Product, error) {
	return products.Product{
		Name:     "club colombia dorada",
		Category: "alcohol",
		Type:     "beer",
	}, nil
}

func (manager *ProductManagerImpl) Delete(productId string) (products.Product, error) {
	return products.Product{
		Name:     "club colombia roja",
		Category: "alcohol",
		Type:     "beer",
	}, nil
}

func (manager *ProductManagerImpl) Create(product products.Product) (products.Product, error) {
	return products.Product{
		Name:     "delirium nocturnum",
		Category: "alcohol",
		Type:     "beer",
	}, nil
}

func (manager *ProductManagerImpl) Update(product products.Product) (products.Product, error) {
	return products.Product{
		Name:     "delirium tremens",
		Category: "alcohol",
		Type:     "beer",
	}, nil
}
