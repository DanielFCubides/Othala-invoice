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

func (manager *ProductManagerImpl) GetAll() ([]*products.Product, error) {
	all, err := manager.reader.FindAll([]string{})
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (manager *ProductManagerImpl) GetById(productId string) (*products.Product, error) {
	product, err := manager.reader.FindById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (manager *ProductManagerImpl) Delete(productName string) (*products.Product, error) {
	product, err := manager.writer.Delete(productName)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (manager *ProductManagerImpl) Create(p products.Product) (*products.Product, error) {
	product, err := manager.writer.Create(p)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (manager *ProductManagerImpl) Update(p products.Product) (*products.Product, error) {
	product, err := manager.writer.Update(p)
	if err != nil {
		return nil, err
	}
	return product, err
}
