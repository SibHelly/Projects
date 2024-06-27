package routes

import (
	"products/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterProductsStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/products/", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products/{productId}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/products/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{productId}", controllers.DeleteProduct).Methods("DELETE")
}
