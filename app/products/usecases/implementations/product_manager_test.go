package implementations

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"othala/app/products"
	"othala/app/products/repositories/mocks"
	"othala/app/products/usecases"
	"testing"
)

type ProductManagerSuite struct {
	suite.Suite
	repoReader *mocks.ProductReader
	repoWriter *mocks.ProductWriter
	manager    usecases.ProductManager
}

func TestRestInit(t *testing.T) {
	suite.Run(t, new(ProductManagerSuite))
}

func (s *ProductManagerSuite) SetupSuite() {
	s.repoReader = new(mocks.ProductReader)
	s.repoWriter = new(mocks.ProductWriter)
	s.manager = newProductManagerImpl(s.repoReader, s.repoWriter)
}

func (s *ProductManagerSuite) TestProductManager_GetById() {
	product, err := s.manager.GetById("milk")
	if err != nil {
		s.Fail(err.Error())
		return

	}
	s.Assert().Equal("milk", product.Name)
	s.Assert().Equal("beer", product.Type)
	s.Assert().Equal("alcohol", product.Category)
}

func (s *ProductManagerSuite) TestProductManager_GetAll() {

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
	s.repoWriter.Mock.On("Delete",mock.Anything).Return(&expected, nil)
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
	s.repoWriter.Mock.On("Create",mock.Anything).Return(&expected, nil)
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
	s.repoWriter.Mock.On("Update",mock.Anything).Return(&expected, nil)
	created, err := s.manager.Update(expected)
	s.Assert().NoError(err)
	s.Assert().NotNil(created)
	s.Assert().Equal(name, created.Name)
	s.Assert().NotNil(category, created.Category)
	s.Assert().NotNil(productType, created.Type)
}
