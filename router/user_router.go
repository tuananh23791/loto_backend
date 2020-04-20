package router

import (
	"travel/controller"

	"github.com/gorilla/mux"
)

func initialRouterUser(router *mux.Router) {
	router.HandleFunc("/auth/signup", controller.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/update", controller.UpdateHandler).Methods("POST")
}
