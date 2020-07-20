package deliveries

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"othala/app/config"
	"othala/app/products/adapters"
)

func Run() {
	err := config.Injector.Invoke(
		func(productAdapter *adapters.RestAdapter) error {
			router := chi.NewRouter()
			port := os.Getenv("PORT")
			registerCommonRoutes(router)
			router.Route("/v1", func(apiRouter chi.Router) {
				adapters.RegisterProductRoutes(apiRouter, productAdapter)
			})
			log.Println("server running on port: " + port)
			return http.ListenAndServe(":"+port, router)
		})
	if err != nil {
		log.Println("error invoking startup code", err)
	}
}

func registerCommonRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("othala OK"))
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("othala OK"))
	})
}
