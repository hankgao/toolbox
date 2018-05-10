package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CsrfValue is something
type CsrfValue struct {
	CsrfToken string `json:"csrf_token"`
}

const (
	maximumSeq       = 3954
	webInterfacePort = 8421
	urlWalletCreate  = "http://127.0.0.1:8421/wallet/create"
	urlCsrf          = "http://127.0.0.1:8421/csrf"
	urlWalletBalance = "http://127.0.0.1:8421/wallet/balance"
)

type addrItem struct {
	addr    string
	balance float64
}

func getWalletBalance(wn string) (string, string) {
	resp, err := http.Get(fmt.Sprintf("%s?id=%s", urlWalletBalance, wn))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)

	r.ReadString('\n')
	r.ReadString('\n')
	c, _ := r.ReadString('\n')
	h, _ := r.ReadString('\n')

	coins := strings.Split(c, ":")[1]
	coins = coins[0 : len(coins)-2]
	hours := strings.Split(h, ":")[1]
	hours = hours[1 : len(hours)-1]

	fmt.Printf("%s#%s\n", coins, hours)

	return coins, hours
}

func main() {

	getWalletBalance("distribution002.wlt")

	// book, accounts := readBook("suncoin.airdrop.csv")

	// source pointer, destionation pointer
	// sp, dp := 0, 0
	// currentWallet := fmt.Sprintf("distribution%03d.wlt", sp)

	// wb, err := getWalletBalance(currentWallet)

	// check next distribution

}

// CreateDistributionWallets does someting
// func CreateDistributionWallets() {
// 	seeds := read100Seeds("suncoin2.100.seeds.csv")
// 	createDistributionWallets(seeds)
// }

// ScanSunAddresses get all Suncoin addresses
// func ScanSunAddresses() {
// 	uxpool := getUxPool()
// 	addressBook := aggregateAddresses(uxpool)

// 	i := 0
// 	for addr, coins := range addressBook {
// 		fmt.Printf("%03d,%s,%f\n", i, addr, coins)
// 		i++
// 	}
// }

func getCsrfToken() string {
	csrf := CsrfValue{}
	resp, err := http.Get(urlCsrf)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(data, &csrf)

	return csrf.CsrfToken

}

func getNextTargets(wb float64, book []addrItem) []addrItem {

	// book[0].balance

	return []addrItem{}
}

func airDoroSun2(book []addrItem) error {

	for _, item := range book {
		fmt.Println(item)
	}

	return nil
}
