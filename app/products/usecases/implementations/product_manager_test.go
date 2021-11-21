package implementations

import (
	"github.com/stretchr/testify/suite"
	"othala/app/products/usecases"
	"othala/app/products/usecases/mocks"
	"testing"
)

type ProductManagerSuite struct {
	suite.Suite
	repository *mocks.ProductManager
	manager    usecases.ProductManager
}

func TestRestInit(t *testing.T) {
	suite.Run(t, new(ProductManagerSuite))
}

func (s *ProductManagerSuite) SetupSuite() {
	s.repository = new(mocks.ProductManager)
	s.manager = newProductManagerImpl()
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

}
func (s *ProductManagerSuite) TestProductManager_Create() {

}
func (s *ProductManagerSuite) TestProductManager_Update() {

}
