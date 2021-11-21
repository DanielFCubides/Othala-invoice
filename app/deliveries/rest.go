package deliveries

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
	"othala/app/config"
	"othala/app/products/adapters"
)

func Run() {
	err := config.Injector.Invoke(AddCorsHeaders)
	if err != nil {
		log.Println("error invoking startup code", err)
	}
}

func AddCorsHeaders(productAdapter *adapters.RestAdapter) error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		//AllowCredentials: false,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	port := os.Getenv("PORT")
	registerCommonRoutes(router)
	router.Route("/v1", func(apiRouter chi.Router) {
		adapters.RegisterProductRoutes(apiRouter, productAdapter)
	})
	log.Println("server running on port: " + port)
	return http.ListenAndServe(":"+port, router)
}

func registerCommonRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("othala OK"))
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("othala OK"))
	})
}
