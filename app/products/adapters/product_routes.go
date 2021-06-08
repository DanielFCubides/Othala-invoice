package adapters

import "github.com/go-chi/chi"

func RegisterProductRoutes(r chi.Router, adapter *RestAdapter) {
	r.Get("/products/", adapter.GetProducts)
	r.Get("/products/{productId}/", adapter.GetProductById)
	r.Post("/products/", adapter.CreateProduct)
	r.Put("/products/{productId}/", adapter.UpdateProduct)
	r.Delete("/products/{productId}/", adapter.DeleteProduct)
}
