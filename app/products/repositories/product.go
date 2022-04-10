package repositories

import (
	"othala/app/config"
	"othala/app/products"
)

type Product struct {
	Name     string `gorm:"primaryKey"`
	Category string
	Type     string
}

func (p Product) ToDomain() products.Product {
	return products.Product{
		Name:     p.Name,
		Category: p.Category,
		Type:     p.Type,
	}
}

func (p Product) FromDomain(product products.Product) Product {
	return Product{
		Name:     product.Name,
		Category: product.Category,
		Type:     product.Type,
	}
}

func Migrate()  {
	var conn config.Connection
	invokeFunc := func(connection config.Connection) {
		conn = connection
	}
	err := config.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	db := conn.GetDatabase()

	err = db.AutoMigrate(Product{})
	if err != nil {
		panic(err)
	}

}