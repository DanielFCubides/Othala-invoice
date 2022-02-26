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

func (p ProductWriter) Create(product products.Product) (*products.Product, error) {
	panic("implement me")
}

func (p ProductWriter) CreateBatch(products []products.Product) (*products.Product, error) {
	panic("implement me")
}

func (p ProductWriter) Update(product products.Product) (*products.Product, error) {
	panic("implement me")
}

func (p ProductWriter) Delete(product products.Product) (*products.Product, error) {
	panic("implement me")
}
