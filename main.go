package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mistikel/inventoy/api"
)

func main() {
	router := httprouter.New()

	router.GET("/", welcome)

	api_v1 := api.Api()
	api_v1.Register(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome to inventory"))
}
