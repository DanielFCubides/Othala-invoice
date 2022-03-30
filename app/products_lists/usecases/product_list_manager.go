package usecases

import "othala/app/products_lists"

type ProductListManager interface {
	Create(lists products_lists.ProductLists) (*products_lists.ProductLists, error)
	Update(ID int, lists products_lists.ProductLists) (*products_lists.ProductLists, error)
	Delete(ID int) (*products_lists.ProductLists, error)
	GetByID(ID int) (*products_lists.ProductLists, error)
}
