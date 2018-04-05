package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// 1.a set routing dari item
	router = setItemRouters(router)
	return router
}
