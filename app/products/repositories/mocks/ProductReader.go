// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	products "othala/app/products"

	mock "github.com/stretchr/testify/mock"
)

// ProductReader is an autogenerated mock type for the ProductReader type
type ProductReader struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: searchParams
func (_m *ProductReader) FindAll(searchParams []string) ([]*products.Product, error) {
	ret := _m.Called(searchParams)

	var r0 []*products.Product
	if rf, ok := ret.Get(0).(func([]string) []*products.Product); ok {
		r0 = rf(searchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*products.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(searchParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *ProductReader) FindById(id string) (*products.Product, error) {
	ret := _m.Called(id)

	var r0 *products.Product
	if rf, ok := ret.Get(0).(func(string) *products.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
