package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vaclav-dvorak/go-api-hello-world/api"
	_ "github.com/vaclav-dvorak/go-api-hello-world/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", api.HandleHello).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
