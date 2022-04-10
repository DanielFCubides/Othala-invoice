package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"othala/app/config"
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
	response, err := adapter.manager.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var productRequest ProductRequest
	err := json.NewDecoder(request.Body).Decode(&productRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	product := productRequest.mapToProduct()
	response, _ := adapter.manager.Update(product)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) GetProductById(writer http.ResponseWriter, request *http.Request) {
	productID := chi.URLParam(request,"productId")
	if len(productID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(fmt.Sprint("{'error':'no product id at the request'}"))
		return
	}
	response, err := adapter.manager.GetById(productID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var productRequest ProductRequest
	err := json.NewDecoder(request.Body).Decode(&productRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response, _ := adapter.manager.Create(productRequest.mapToProduct())
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}

func (adapter *RestAdapter) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	productID := chi.URLParam(request,"productId")
	if len(productID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(fmt.Sprint("{'error':'no product id at the request'}"))
		return
	}
	response, _ := adapter.manager.Delete(productID)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(response)
}
