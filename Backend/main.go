package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	historialService "sopes-practica1/service"
)

func ReturnReport(w http.ResponseWriter, req *http.Request) {
	historial, err := historialService.Read()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(historial)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/oper/report", ReturnReport).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
}
