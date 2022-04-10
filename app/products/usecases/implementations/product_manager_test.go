package implementations

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"othala/app/products"
	"othala/app/products/repositories/mocks"
	"othala/app/products/usecases"
	"testing"
)

type ProductManagerSuite struct {
	suite.Suite
	reader  *mocks.ProductReader
	writer  *mocks.ProductWriter
	manager usecases.ProductManager
}

func TestRestInit(t *testing.T) {
	suite.Run(t, new(ProductManagerSuite))
}

func (s *ProductManagerSuite) SetupSuite() {
	s.reader = new(mocks.ProductReader)
	s.writer = new(mocks.ProductWriter)
	s.manager = newProductManagerImpl(s.reader, s.writer)
}

func (s *ProductManagerSuite) TestProductManager_GetById_Success() {
	arguments := &products.Product{
		Name:     "twix",
		Category: "chocolate",
		Type:     "candy",
	}
	s.reader.On("FindById", mock.Anything).Return(arguments, nil)
	product, err := s.manager.GetById("milk")
	if err != nil {
		s.Fail(err.Error())
		return
	}
	s.Assert().Equal("twix", product.Name)
	s.Assert().Equal("candy", product.Type)
	s.Assert().Equal("chocolate", product.Category)
}

func (s *ProductManagerSuite) TestProductManager_GetById_Failed() {
	s.reader.On("FindById", mock.Anything).Return(nil, errors.New("some error"))
	_, err := s.manager.GetById("milk")
	if err != nil {
		s.Error(err)
		return
	}
}

func (s *ProductManagerSuite) TestProductManager_Delete() {
	name := "ron viejo de caldas"
	category := "alcohol"
	productType := "rum"
	expected := products.Product{
		Name:     name,
		Category: category,
		Type:     productType,
	}
	s.writer.Mock.On("Delete",mock.Anything).Return(&expected, nil)
	created, err := s.manager.Delete("ron viejo de caldas")
	s.Assert().NoError(err)
	s.Assert().NotNil(created)
	s.Assert().Equal(name, created.Name)
	s.Assert().NotNil(category, created.Category)
	s.Assert().NotNil(productType, created.Type)
}
func (s *ProductManagerSuite) TestProductManager_Create() {
	name := "ron viejo de caldas"
	category := "alcohol"
	productType := "rum"
	expected := products.Product{
		Name:     name,
		Category: category,
		Type:     productType,
	}
	s.writer.Mock.On("Create",mock.Anything).Return(&expected, nil)
	created, err := s.manager.Create(expected)
	s.Assert().NoError(err)
	s.Assert().NotNil(created)
	s.Assert().Equal(name, created.Name)
	s.Assert().NotNil(category, created.Category)
	s.Assert().NotNil(productType, created.Type)

}
func (s *ProductManagerSuite) TestProductManager_Update() {
	name := "ron viejo de caldas"
	category := "alcohol"
	productType := "rum"
	expected := products.Product{
		Name:     name,
		Category: category,
		Type:     productType,
	}
	s.writer.Mock.On("Update",mock.Anything).Return(&expected, nil)
	created, err := s.manager.Update(expected)
	s.Assert().NoError(err)
	s.Assert().NotNil(created)
	s.Assert().Equal(name, created.Name)
	s.Assert().NotNil(category, created.Category)
	s.Assert().NotNil(productType, created.Type)
}

