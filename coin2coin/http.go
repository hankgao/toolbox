package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func run() {
	r := mux.NewRouter()
	r.HandleFunc("/api/coinpair", coinPairHandler).Methods("GET")
	r.HandleFunc("/api/order", orderHandler).Methods("PUT")
	r.HandleFunc("/", homeHandler).Methods("GET")
}

func coinPairHandler(w http.ResponseWriter, r *http.Request) {
	pairs := LoadCoinPairs()

	cpr := CoinPairsR{
		Number:    len(pairs),
		CoinPairs: pairs,
	}

	mcpr, _ := json.MarshalIndent(cpr, "", "\t")

	//TODO tell cliemt is a json formated response

	w.Write(mcpr)

}

// orderHandler handles PUT /api/order
// it expects the following JSON data in the request body
// {
// 	"coin_from": "mzc",
// 	"coin_to": "shellcoin"
// 	"receiving_address": "23skfksfkfksdfjasdfweqrwexf" // given by investor
// 	"despoit_address":  // is this property necessary?
// 	"investor_name": "Hank Gao"
// 	"investor_mobile": "18618571330"
// 	"investor_address": "上海市浦东光辉路23弄2号220"
// 	"time_placed": "2018-09-10 15:01:02"
// }
//
//
//
func orderHandler(w http.ResponseWriter, r *http.Request) {
	order := Order{}
	order.CoinFrom = ""
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO: respond gracefully
		panic(err)
	}

	x := string(b)
	fmt.Println(x)

	err = json.Unmarshal(b, &order)
	if err != nil {
		// TODO: respond gracefully
		panic(err)
	}

	if isValidOrder(order) == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid request:\n%s", b)))
		return
	}

	// Now start a transaction
	// The seed we use:
	// helmet amazing fetch month sunset session frame output tackle immune enhance simple
	// Generate a new address
	// update database
	_, err = db.stmtOrders.Exec(order.CoinFrom, order.CoinTo, order.DepositingAddr, order.ReceivingAddr, order.InvestorName,
		order.InvestorMobile, order.InvestorAddress, order.TimePlaced)

	if err != nil {
		// TODO: respond gracefully
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create order in database"))
		return
	}

	// finally respond
	w.Write([]byte("This is a test"))

}

func isValidOrder(o Order) bool {
	// check if investor provided receiving coin address
	if o.ReceivingAddr != "" {
		return true
	}

	// if not, then check mailing address is provided or not
	if o.InvestorName != "" && o.InvestorMobile != "" && o.InvestorAddress != "" {
		return true
	}

	return false
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}
