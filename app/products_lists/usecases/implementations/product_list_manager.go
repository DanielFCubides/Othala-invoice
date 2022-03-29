package implementations

import (
	"errors"
	"othala/app/config"
	"othala/app/products_lists"
	"othala/app/products_lists/usecases"
)

type ProductListManager struct {
}

func (p ProductListManager) Create(lists products_lists.ProductLists) (*products_lists.ProductLists, error) {
	if len(lists.Products) == 0 {
		return nil, errors.New("empty list")
	}
	return &lists, nil
}

func (p ProductListManager) Update(ID int, lists products_lists.ProductLists) (*products_lists.ProductLists, error) {
	return &lists, nil
}

func (p ProductListManager) Delete(ID int) (*products_lists.ProductLists, error) {
	return nil, nil
}

func (p ProductListManager) GetByID(ID int) (*products_lists.ProductLists, error) {
	return nil, nil
}

func init() {
	if err := config.Injector.Provide(newProductListsManagerImpl); err != nil {
		panic("error providing ProductListManager")
	}
}

func newProductListsManagerImpl() usecases.ProductListManager {
	return &ProductListManager{}
}
