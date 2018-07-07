package skytb

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
)

// SpendFromUx spends coins from a specified ux, allowing multiple destinations
func SpendFromUx(uxid string, dst []Balance) error {
	return nil
}

// SpendFromAddress spends coins from a specified address
func SpendFromAddress(addr string, dst []Balance) error {
	return nil
}

// SpendCoinFromWallet sends a certain amout of coins from a given wallet of ct coin type
func SpendCoinFromWallet(wn, cn string, target AddrItem) {

	coins := strconv.FormatUint(target.Balance, 10)

	v := url.Values{}

	v.Set("id", wn)
	v.Set("dst", target.Addr)
	v.Set("coins", coins)

	body := v.Encode()

	ctd := CoinTypeDetails(cn)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%4d/wallet/spend", ctd.WebInterfacePort), strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-CSRF-Token", getCsrfToken(cn))

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(resp.Status)
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

// DistributionTransaction returns a transaction in which the coins in a utxo is distributed to specified addresses
// uxid   -
// seckey - private key
// addrs  - addressed to be distributed into
// each   - number of coins to be distributed into each address
func DistributionTransaction(uxid, seckey string, addrs []string, each uint64) coin.Transaction {
	var tx coin.Transaction

	output := cipher.MustSHA256FromHex(uxid)
	tx.PushInput(output)

	for i := range addrs {
		addr := cipher.MustDecodeBase58Address(addrs[i])
		tx.PushOutput(addr, each*1e6, each)
	}

	seckeys := make([]cipher.SecKey, 1)
	seckeys[0] = cipher.MustSecKeyFromHex(seckey)
	tx.SignInputs(seckeys)

	tx.UpdateHeader()

	err := tx.Verify()

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("signature= %s", tx.Sigs[0].Hex())

	return tx
}

// DistCoinsNamedWallets distributes coins from 100 named wallets. Wallets should be named in the following format:
//  distribution%%%.wlt, where %%% is from 0 to 099, note the leading zero(s)
func DistCoinsNamedWallets(cn string, targets []AddrItem, count int) {
	wltIndex := 0
	for i, target := range targets {
		if i >= count-1 {
			break
		}

		spent := false
		for wi := 0; wi < 100; wi++ {
			wltName := fmt.Sprintf("distribution%03d.wlt", (wltIndex+wi)%100)

			coins, _ := getWalletBalance(wltName, cn)

			if coins >= target.Balance {
				fmt.Printf("Distribute %d droplets to %s from %s\n", target.Balance, target.Addr, wltName)

				SpendCoinFromWallet(wltName, cn, target)

				wltIndex = wltIndex + wi
				spent = true

				time.Sleep(2 * time.Second)

				break

			}

			wi++
		}

		if !spent {
			fmt.Printf("Warning: failed to airdrop %d droplets to address =>%s\n", target.Balance, target.Addr)
		}

		wltIndex++
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
