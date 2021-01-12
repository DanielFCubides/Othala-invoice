package implementations

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	panic("implement me")
}

func (p ProductReader) FindAll(searchParams []string) ([]*products.Product, error) {
	panic("implement me")
}
