package adapters

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
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
	r.useCase.Mock.On("GetAll").Return([]*products.Product{{
		Name:     productName,
		Category: "lactose",
		Type:     "freezer",
	}}, nil)

	recorder, err := r.callAdapter(http.MethodGet, "/v1/products/", nil, r.adapter.GetProducts)
	if recorder == nil {
		r.Fail("could not call adapter")
		return
	}
	r.Equal(http.StatusOK, recorder.Code)
	var response []*ProductResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		r.Fail("could not parse response")
		return
	}
	r.Equal(productName, response[0].Name)
}

func (r RestSuite) TestRest_GetProductByID() {
	productName := "Milk"
	r.useCase.Mock.On("GetById", mock.Anything).
		Return(&products.Product{
			Name:     productName,
			Category: "lactose",
			Type:     "freezer",
		}, nil)

	recorder, err := r.callAdapter(http.MethodGet, "/v1/products/1/", nil, r.adapter.GetProductById)
	if recorder == nil {
		r.Fail("could not call adapter")
		return
	}
	r.Equal(http.StatusOK, recorder.Code)
	var response *ProductResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		r.Fail("could not parse response")
		return
	}
	r.Equal(productName, response.Name)
}

func (r RestSuite) TestRest_CreateProduct() {
	productName := "Milk"
	product := products.Product{
		Name:     productName,
		Category: "lactose",
		Type:     "freezer",
	}
	r.useCase.Mock.On("Create", mock.Anything).
		Return(&product, nil)
	b, err := json.Marshal(product)
	reader := bytes.NewBuffer(b)
	recorder, err := r.callAdapter(http.MethodPost, "/v1/products/", reader, r.adapter.CreateProduct)
	if recorder == nil {
		r.Fail("could not call adapter")
		return
	}
	r.Equal(http.StatusOK, recorder.Code)
	var response *ProductResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		r.Fail("could not parse response")
		return
	}
	r.Equal(productName, response.Name)
}

func (r RestSuite) TestRest_UpdateProduct() {
	productName := "Milk"
	product := products.Product{
		Name:     productName,
		Category: "lactose",
		Type:     "freezer",
	}
	r.useCase.Mock.On("Update", mock.Anything).
		Return(&product, nil)
	b, err := json.Marshal(product)
	reader := bytes.NewBuffer(b)
	recorder, err := r.callAdapter(http.MethodPut, "/v1/products/1/", reader, r.adapter.UpdateProduct)
	if recorder == nil {
		r.Fail("could not call adapter")
		return
	}
	r.Equal(http.StatusOK, recorder.Code)
	var response *ProductResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		r.Fail("could not parse response")
		return
	}
	r.Equal(productName, response.Name)
}

func (r RestSuite) callAdapter(method, url string, body io.Reader, adapter http.HandlerFunc) (*httptest.ResponseRecorder, error) {
	res, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	router := chi.NewRouter()
	router.Route("/v1", func(apiRouter chi.Router) {
		RegisterProductRoutes(apiRouter, r.adapter)
	})
	recorder := httptest.NewRecorder()
	handler := router
	handler.ServeHTTP(recorder, res)
	return recorder, err
}
