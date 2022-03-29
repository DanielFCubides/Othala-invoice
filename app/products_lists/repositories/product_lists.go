package repositories

import "othala/app/products_lists"

type ProductListsRepository interface {
	Read(ID int)(*products_lists.ProductLists, error)
	Write(list *products_lists.ProductLists)(*products_lists.ProductLists, error)
}