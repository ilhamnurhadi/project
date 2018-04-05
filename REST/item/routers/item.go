package routers

import (
	"day15/item/controllers"

	"github.com/gorilla/mux"
)

func setItemRouters(router *mux.Router) *mux.Router {
	// 1.b buat fungsi di controllers
	router.HandleFunc("/item", controllers.GetItem).Methods("GET")
	return router
}
