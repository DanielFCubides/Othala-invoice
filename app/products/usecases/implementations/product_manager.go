package implementations

import (
	"othala/app/config"
	"othala/app/products"
	"othala/app/products/repositories"
	"othala/app/products/usecases"
)

type ProductManagerImpl struct {
	reader repositories.ProductReader
	writer repositories.ProductWriter
}

func init() {
	if err := config.Injector.Provide(newProductManagerImpl); err != nil {
		panic("error providing product manager")
	}
}

func newProductManagerImpl(reader repositories.ProductReader, writer repositories.ProductWriter) usecases.ProductManager {
	return &ProductManagerImpl{
		reader: reader,
		writer: writer,
	}
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
		Name:     productId,
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
