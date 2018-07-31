package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is coin2coin server")
	db.Init()
	defer db.Close()

	startServer()

	pairs := LoadCoinPairs()

	cpr := CoinPairsR{}

	cpr.Number = len(pairs)
	cpr.CoinPairs = pairs

	mcpr, err := json.MarshalIndent(cpr, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(mcpr))

}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/coinpair", coinPairHandler).Methods("GET")
	r.HandleFunc("/api/order", orderHandler).Methods("POST")
	r.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", r)
}
