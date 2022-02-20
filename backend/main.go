package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	m "sopes-practica1/models"
	historialService "sopes-practica1/service"

	"github.com/gorilla/mux"
)

func Operarar(n1 float64, n2 float64, op string) float64 {
	switch op {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		if n2 == 0 {
			return 0
		}
		return n1 / n2
	default:
		return 0
	}
}

func ExecuteOper(w http.ResponseWriter, req *http.Request) {
	var historial m.Historial
	_ = json.NewDecoder(req.Body).Decode(&historial)

	historial.Resultado = Operarar(historial.Numero1, historial.Numero2, historial.Operacion)
	historial.Fecha = time.Now()

	err := historialService.Create(historial)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(historial)
}

func ReturnReport(w http.ResponseWriter, req *http.Request) {
	historial, err := historialService.Read()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(historial)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/oper", ExecuteOper).Methods("POST")
	router.HandleFunc("/oper/report", ReturnReport).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
}
