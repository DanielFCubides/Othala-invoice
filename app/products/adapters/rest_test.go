package adapters

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"othala/app/products"
	"othala/app/products/usecases/mocks"
	"testing"
)

type RestSuite struct {
	suite.Suite
	useCase *mocks.ProductManager
	adapter *RestAdapter
}

func TestRestInit(t *testing.T) {
	suite.Run(t, new(RestSuite))
}

func (r *RestSuite) SetupSuite() {
	r.useCase = new(mocks.ProductManager)
	r.adapter = newRestAdapter(r.useCase)
}

func (r RestSuite) TestRest_GetProducts() {
	productName := "Milk"
	r.useCase.Mock.On("GetAll").Return([]products.Product{{
		Name:     productName,
		Category: "lactose",
		Type:     "freezer",
	}}, nil)
	res, err := http.NewRequest(http.MethodGet, "/products/", nil)

	if err != nil {
		r.Fail("could not create request")
	}

	recoder := httptest.NewRecorder()
	handler := http.HandlerFunc(r.adapter.GetProducts)
	handler.ServeHTTP(recoder, res)

	if status := recoder.Code; status != http.StatusOK {
		r.Fail("error not ok")
	}

	var response []*ProductResponse
	err = json.NewDecoder(recoder.Body).Decode(&response)
	if err != nil {
		r.Fail("could not parse response")
		return
	}
	r.Equal(productName, response[0].Name)

}
