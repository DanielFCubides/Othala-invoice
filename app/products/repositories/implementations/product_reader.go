package implementations

import (
	"fmt"
	"gorm.io/gorm"
	"othala/app/config"
	"othala/app/products"
	"othala/app/products/repositories"
)

type ProductReader struct {
	db *gorm.DB
}

func newProductReader(connection config.Connection) repositories.ProductReader {
	return &ProductReader{db: connection.GetDatabase()}
}

func init() {
	err := config.Injector.Provide(newProductReader)
	if err != nil {
		fmt.Println("Error providing the ProductReader:", err)
		panic(err)
	}
}

func (p ProductReader) FindById(id string) (*products.Product, error) {
	product := repositories.Product{}
	err := p.db.First(&product, "name = ?", id).Error
	if err != nil {
		return nil, err
	}
	domain := product.ToDomain()
	return &domain, nil
}

func (p ProductReader) FindAll(searchParams []string) ([]*products.Product, error) {
	productsList := make([]*repositories.Product, 0)
	output := make([]*products.Product, 0)
	err := p.db.Find(&productsList).Error
	if err != nil {
		return nil, err
	}
	for i, a := range productsList {
		print(i, a)
		domain := a.ToDomain()
		output = append(output, &domain)
	}
	return output, nil
}
