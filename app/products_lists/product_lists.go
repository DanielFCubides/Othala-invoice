package products_lists

import "othala/app/products"

type Items struct {
	Product products.Product
	Price float64
}

type ProductLists struct {
	ID int
	Products []Items
}

func (pl *ProductLists) Price()  {
	result := 0.0
	for _, v := range pl.Products {
		result += v.Price
	}
}