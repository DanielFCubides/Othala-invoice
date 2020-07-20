package adapters

import (
	"encoding/json"
	"net/http"
	"othala/app/config"
)

type RestAdapter struct {
}

func init() {
	err := config.Injector.Provide(newRestAdapter)
	if err != nil {
		panic("Error providing products rest adapter")
	}
}

func newRestAdapter() *RestAdapter {
	return &RestAdapter{}
}

func (adapter *RestAdapter) GetProducts(writer http.ResponseWriter, r *http.Request) {
	response := []ProductResponse{{
		Name:     "Club Colombia Negra",
		Category: "Alcohol",
		Type:     "Beer",
	}}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	_ = json.NewEncoder(writer).Encode(response)
}
