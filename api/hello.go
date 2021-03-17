package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HandleHello(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Println("Endpoint Hit: hello")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HelloResponse{Message: "Hello " + name})
}
