package adapters

import (
	"encoding/json"
	"net/http"
	"othala/app/config"
	"othala/app/products"
	"othala/app/products/usecases"
)

type RestAdapter struct {
	manager usecases.ProductManager
}

func init() {
	err := config.Injector.Provide(newRestAdapter)
	if err != nil {
		panic("Error providing products rest adapter")
	}
}

func newRestAdapter(manager usecases.ProductManager) *RestAdapter {
	return &RestAdapter{
		manager: manager,
	}
}

func (adapter *RestAdapter) GetProducts(writer http.ResponseWriter, r *http.Request) {
	response, _ := adapter.manager.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	product := products.Product{}
	response, _ := adapter.manager.Update(product)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) GetProductById(writer http.ResponseWriter, request *http.Request) {
	response, _ := adapter.manager.GetById("product")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	product := products.Product{}
	response, _ := adapter.manager.Create(product)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	response, _ := adapter.manager.Delete("product")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}
