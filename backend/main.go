package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	m "sopes-practica1/models"
	historialService "sopes-practica1/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Oper struct {
	Numero1   string `json:"numero1,omitempty"`
	Numero2   string `json:"numero2,omitempty"`
	Operacion string `json:"operacion,omitempty"`
}

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
	var op Oper
	_ = json.NewDecoder(req.Body).Decode(&op)

	n1, er := strconv.ParseFloat(op.Numero1, 64)
	n2, er := strconv.ParseFloat(op.Numero2, 64)
	if er != nil {
		json.NewEncoder(w).Encode(er)
	}
	historial.Numero1 = n1
	historial.Numero2 = n2
	historial.Operacion = op.Operacion
	historial.Resultado = Operarar(historial.Numero1, historial.Numero2, historial.Operacion)
	historial.Fecha = time.Now()

	err := historialService.Create(historial)

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(historial)
}

func ReturnReport(w http.ResponseWriter, req *http.Request) {

	historial, err := historialService.Read()

	historial = ordenarMenor(historial, len(historial))

	w.Header().Set("Content-type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(historial)
}

func ordenarMenor(listHiss m.Historiales, Cant int) m.Historiales {
	var tmp *m.Historial

	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listHiss[x].Fecha.After(listHiss[y].Fecha) {
				tmp = listHiss[y]
				listHiss[y] = listHiss[x]
				listHiss[x] = tmp
			}
		}
	}

	return listHiss
}

func main() {
	router := mux.NewRouter()

	heaers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/oper", ExecuteOper).Methods("POST")
	router.HandleFunc("/oper/report", ReturnReport).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(heaers, credentials, methods, ttl, origins)(router)))
}
