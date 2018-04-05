package main

import (
	"day15/item/routers"
	"log"
	"net/http"
)

func main() {
	// p1. buat routing
	// buat fungsi routing
	router := routers.InitRouters()

	//buatlah configurasi server
	log.Fatal(http.ListenAndServe(":8886", router))
}
