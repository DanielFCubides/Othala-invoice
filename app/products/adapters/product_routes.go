package adapters

import "github.com/go-chi/chi"

func RegisterProductRoutes(r chi.Router, adapter *RestAdapter) {
	r.Get("/products", adapter.GetProducts)
}
