package implementations

import (
	"fmt"
	"gorm.io/gorm"
	"othala/app/config"
	"othala/app/products"
	"othala/app/products/repositories"
)

type ProductWriter struct {
	db *gorm.DB
}

func NewProductWriter(connection config.Connection) repositories.ProductWriter {
	return &ProductWriter{db: connection.GetDatabase()}
}

func init() {
	err := config.Injector.Provide(NewProductWriter)
	if err != nil {
		fmt.Println("Error providing the ProductWriter:", err)
		panic(err)
	}
}

func (w ProductWriter) Create(p products.Product) (*products.Product, error) {
	product := repositories.Product{}.FromDomain(p)
	result := w.db.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	domain := product.ToDomain()
	return &domain, nil
}

func (w ProductWriter) CreateBatch(products []products.Product) (*products.Product, error) {
	panic("implement me")
}

func (w ProductWriter) Update(p products.Product) (*products.Product, error) {
	product := repositories.Product{}.FromDomain(p)
	created := w.db.First(repositories.Product{}, p.Name)
	if created.Error != nil {
		return nil, created.Error
	}
	result := w.db.Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	domain := product.ToDomain()
	return &domain, nil
}

func (w ProductWriter) Delete(productName string) (*products.Product, error) {

	product := repositories.Product{}
	result := w.db.First(product, productName)
	if result.Error != nil {
		return nil, result.Error
	}
	result = w.db.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	domain := product.ToDomain()
	return &domain, nil
}
