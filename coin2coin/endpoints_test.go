package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPairRequest(t *testing.T) {

	db.Init()
	defer db.Close()

	// create sample request
	req, err := http.NewRequest("GET", "/api/coinpair", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(coinPairHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestOrderRequest(t *testing.T) {

	db.Init()
	defer db.Close()

	// create sample request
	o := Order{
		CoinFrom:        "mzcoin",
		CoinTo:          "shellcoin",
		ReceivingAddr:   "",
		DepositingAddr:  "",
		InvestorName:    "Hank Gao",
		InvestorMobile:  "1873453454",
		InvestorAddress: "上海市浦东新区光晖路130号5弄101",
		TimePlaced:      "2018-09-10 15:30:23",
	}

	mo, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/api/order", bytes.NewBuffer(mo))
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(orderHandler)
	handler.ServeHTTP(rr, req)

	body, _ := ioutil.ReadAll(rr.Body)

	t.Error(string(body))

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestCoinPair(t *testing.T) {

}
