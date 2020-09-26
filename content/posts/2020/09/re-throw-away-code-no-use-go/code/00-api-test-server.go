package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type Stock struct {
	Symbol string  `json:"Symbol"`
	Name   string  `json:"Name"`
	Price  float64 `json:"Price"`
}

func chaosMonkey() {
	max := 1000
	min := 0
	randomNo := rand.Intn(max-min) + min
	if randomNo == 999 {
		log.Fatalf("Oops...")
	}
}

func getPrice(symbol string) (*Stock, int) {
	prices := map[string]*Stock{
		"PEAR":  &Stock{Symbol: "PEAR", Name: "Pear Computers Inc.", Price: 1776.74},
		"SQRL":  &Stock{Symbol: "SQRL", Name: "Squirrel Nut Import/Export LLP", Price: 5.4},
		"NULL":  &Stock{Symbol: "NULL", Name: "Remove this before go-live (Jerry)", Price: .9},
		"ZVZZT": &Stock{Symbol: "ZVZZT", Name: "SuperMontage TEST", Price: 10.0},
	}
	if symbol == "" {
		return nil, 400
	}
	if val, ok := prices[symbol]; ok {
		return val, 200
	}
	return nil, 404
}

func handleStockRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	symbols, ok := r.URL.Query()["symbol"]
	if !ok || len(symbols) <= 0 {
		w.WriteHeader(500)
		return
	}
	// Crash with a 1 in 1000 chance:
	chaosMonkey()

	symbol := symbols[0]
	log.Printf("Symbol requested %s", string(symbol))

	stock, code := getPrice(symbol)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(stock)
}

func main() {
	http.HandleFunc("/stocks", handleStockRequest)
	http.ListenAndServe(":3000", nil)
}
