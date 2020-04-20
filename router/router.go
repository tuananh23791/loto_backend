package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Initial() {
	r := mux.NewRouter()
	fmt.Println("setup router")
	initialRouterUser(r)
	initialRouterCity(r)

	fmt.Println("run success")
	http.ListenAndServe(":3000", r)
}
