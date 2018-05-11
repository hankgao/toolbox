package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	urlSpend         = "http://127.0.0.1:8421/wallet/spend"
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

func spendCoin(wn string, target addrItem) {

	coins := strconv.FormatUint(uint64(target.balance*1000000.0), 10)

	v := url.Values{}

	v.Set("id", wn)
	v.Set("dst", target.addr)
	v.Set("coins", coins)

	body := v.Encode()

	req, err := http.NewRequest("POST", urlSpend, strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-CSRF-Token", getCsrfToken())

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200" {
		panic(err)
	}

	// data, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// // parse result
	// var f interface{}
	// err = json.Unmarshal(data, &f)
	// m := f.(map[string]interface{})

}

func distributeCoins(targets []addrItem) {
	wltIndex := 0

	for _, target := range targets {
		spendCoin(fmt.Sprintf("distribution%03d", wltIndex), target)
		wltIndex = (wltIndex + 1) % 100
	}
}

// func distributeCoins(wlt, change string, targets []addrItem, qty int) error {

// 	jsonBody := `
// 	{
// 		"hours_selection": {
// 			"type": "auto",
// 			"mode": "share",
// 			"share_factor": "0.5"
// 		},
// 		"wallet": {
// 			"id": "walletName"
// 		},
// 		"change_address": "changeAddress",
// 		"to": [
// 		`

// 	// "to": [
// 	// 	{
// 	// 		"address": "targetAddress",
// 	// 		"coins": "targetCoins",
// 	// 	}
// 	// 	]

// 	jsonBody = strings.Replace(jsonBody, "walletName", wlt, 1)
// 	jsonBody = strings.Replace(jsonBody, "changeAddress", change, 1)

// 	to := ""
// 	for i := 0; i < qty; i++ {
// 		to += fmt.Sprintf("{\n\"address\":%s\n", targets[i].addr)
// 		to += fmt.Sprintf("\"address\":%d\n}", uint64(targets[i].balance*1e6))

// 		if i < qty-1 {
// 			to += ","
// 		}

// 		to += "\n"
// 	}

// 	to += "]\n}"

// 	fmt.Println(jsonBody)

// 	req, err := http.NewRequest("POST", urlWalletCreate, strings.NewReader(jsonBody))

// 	if err != nil {
// 		panic(err)
// 	}

// 	req.Header.Add("Content-Type", "application/application/json")
// 	req.Header.Add("X-CSRF-Token", getCsrfToken())

// 	c := &http.Client{}
// 	resp, err := c.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%s\n", string(data))

// 	return nil
// }

func main() {

	readBook("suncoin.airdrop.csv")

	// for i := 0; i < count; i++ {
	// 	// fmt.Printf("%s, %f\n", accounts[i].addr, accounts[i].balance)
	// }

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

func readBook(fn string) ([]addrItem, int) {
	book := make([]addrItem, 2000)
	f, err := os.Open(fn)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		a := strings.Split(line, ",")
		balance, err := strconv.ParseFloat(a[2][0:len(a[2])-2], 64)

		fmt.Printf("%f\n", balance)

		if err != nil {
			panic(err)
		}

		book[count] = addrItem{
			a[1],
			balance,
		}

		count++

	}

	return book, count
}

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
