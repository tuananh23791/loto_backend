package router

import (
	"travel/controller"

	"github.com/gorilla/mux"
)

func initialRouterCity(router *mux.Router) {
	router.HandleFunc("/city/create", controller.CreateCityHandler).Methods("POST")
	router.HandleFunc("/city/get", controller.GetCityHandler).Methods("GET")
}
