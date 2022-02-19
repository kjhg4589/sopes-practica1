package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Operacion struct {
	Num1      float64 `json:"numero1,omitempty"`
	FirstName string  `json:"firstname,omitempty"`
	LastName  string  `json:"lastname,omitempty"`
}

var opers []Operacion

func executeOper(w http.ResponseWriter, req *http.Request) {

}

func returnReport(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(opers)
}

func main() {
	router := mux.NewRouter()

	// adding example data
	opers = append(opers, Operacion{Num1: 1.2, FirstName: "Ryan", LastName: "Ray"})
	opers = append(opers, Operacion{Num1: 2.2, FirstName: "Maria", LastName: "Ray"})

	//endpoints
	router.HandleFunc("/oper", executeOper).Methods("POST")
	router.HandleFunc("/oper/report", returnReport).Methods("GET")

	log.Fatal(http.ListenAndServe(":9080", router))

}
