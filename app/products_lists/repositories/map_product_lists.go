package repositories

import (
	"errors"
	"othala/app/config"
	"othala/app/products_lists"
)

type MapProductListsRepository struct {
	Data map[int]products_lists.ProductLists
}

func init() {
	if err := config.Injector.Provide(newProductListsRepo); err != nil {
		panic("error providing ProductListManager")
	}
}

func newProductListsRepo() ProductListsRepository {
	return &MapProductListsRepository{
		Data: make(map[int]products_lists.ProductLists, 0),
	}
}

func (m MapProductListsRepository) Read(ID int) (*products_lists.ProductLists, error) {
	if value, exist := m.Data[ID]; exist{
		return &value, nil
	}
	return nil, errors.New("element not found")
}

func (m MapProductListsRepository) Write(list *products_lists.ProductLists) (*products_lists.ProductLists, error) {
	if _, exist := m.Data[list.ID]; exist{
		return nil, errors.New("element already exist")
	}
	index := len(m.Data)
	m.Data[index] = *list
	return list, nil
}
