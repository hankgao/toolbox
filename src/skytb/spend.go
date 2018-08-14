package skytb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/skycoin/skycoin/src/api"
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

// DistCoins2OneHundred distributes coins in the genesis block to 100 addresses in addrs
// cn - coin name
// wn - wallet name
func DistCoins2OneHundred(addrs []string, cn, wn string, slice float64) error {
	_, ok := CoinTypesSupported[cn]
	if !ok {
		return fmt.Errorf("%s is not supported", cn)
	}

	body := `
	{
		"hours_selection": {
			"type": "auto",
			"mode": "share",
			"share_factor": "0.5"
		},
		"wallet": {
			"id": "foo.wlt",
			"unspents": ["275f555f7aef20ff30718708f802e30ef36f338d3a0a85d1f61007c6c643a2b3"]
		},
		"to": [
	`
	a2r := struct {
		Address string `json:"address"`
		Coins   string `json:"coins"`
	}{}

	for i := 0; i < 100; i++ {
		a2r.Address = addrs[i]
		a2r.Coins = strconv.FormatFloat(slice, 'f', 3, 64)

		ma2r, err := json.MarshalIndent(a2r, "", "  ")
		if err != nil {
			panic(err)
		}

		body = body + string(ma2r)
		if i < 99 {
			body = body + ",\n"
		}
	}

	body = body + "]}"

	body = strings.Replace(body, "foo.wlt", wn, -1)

	fmt.Println(body)

	// ctd := CoinTypeDetails(cn)
	// req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%d/api/v1/wallet/transaction", ctd.WebInterfacePort), strings.NewReader(body))
	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:49827/api/v1/wallet/transaction"), strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
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

	return nil
}

// DistributeCoins distributes coins to addresses specified in targets
// func DistributeCoins(cn, wn string, targets []AddrItem) error {

// 	_, ok := CoinTypesSupported[cn]
// 	if !ok {
// 		return fmt.Errorf("%s is not supported", cn)
// 	}

// 	// TODO: need to give port
// 	client := api.NewClient("127.0.0.1")

// 	// constrcut transaction request
// 	req := api.CreateTransactionRequest{
// 		HoursSelection: api.HoursSelection{
// 			Type:        "auto",
// 			Mode:        "share",
// 			ShareFactor: "0.1",
// 		},
// 		ChangeAddress:     nil,
// 		IgnoreUnconfirmed: false,
// 		Wallet: api.CreateTransactionRequestWallet{
// 			ID: wn,
// 		},
// 	}

// 	var r api.Receiver
// 	for i := 0; i < len(targets) {
// 		r.Address = targets[i].Addr
// 		r.Coins =
// 		r := api.Receiver {
// 			Address: targets[i].Addr,
// 			Coins: targets[i].Balance,
// 		}
// 	}

// 	body := `
// {
// 	"hours_selection": {
// 		"type": "auto",
// 		"mode": "share",
// 		"share_factor": "0.1"
// 	},
// 	"wallet": {
// 		"id": "foo.wlt"
// 	},
// 	"to": [
// `
// 	a2r := struct {
// 		Address string `json:"address"`
// 		Coins   string `json:"coins"`
// 	}{}

// 	totalCoins := float64(0.00)

// 	// TODO:
// 	// - check whether you have enough coins
// 	// - check if number of targets is bigger than 100
// 	for i := 0; i < len(targets); i++ {
// 		a2r.Address = targets[i].Addr
// 		a2r.Coins = strconv.FormatFloat(float64(targets[i].Balance), 'f', 3, 64)

// 		totalCoins += float64(targets[i].Balance)

// 		ma2r, err := json.MarshalIndent(a2r, "", "  ")
// 		if err != nil {
// 			panic(err)
// 		}

// 		body = body + string(ma2r)
// 		if i < len(targets)-1 {
// 			body = body + ",\n"
// 		}
// 	}

// 	body = body + "]}"

// 	body = strings.Replace(body, "foo.wlt", wn, -1)

// 	fmt.Println("Spending body for checking")
// 	fmt.Println(body)

// 	ctd := CoinTypeDetails(cn)
// 	req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%d/api/v1/wallet/transaction", ctd.WebInterfacePort), strings.NewReader(body))

// 	if err != nil {
// 		panic(err)
// 	}

// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("X-CSRF-Token", getCsrfToken(cn))

// 	c := &http.Client{}
// 	resp, err := c.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	resp.Body.Read()

// 	if resp.StatusCode != 200 {
// 		panic(resp.Status)
// 	} else {
// 		//
// 	}

// 	return nil
// }

// DistributeCoins distributes coins to addresses specified in targets
func DistributeCoins(cn, wn string, targets []api.Receiver) error {

	_, ok := CoinTypesSupported[cn]
	if !ok {
		return fmt.Errorf("%s is not supported", cn)
	}

	// TODO: need to give port
	ctd := CoinTypeDetails(cn)
	addr := fmt.Sprintf("http://127.0.0.1:%d", ctd.WebInterfacePort)
	client := api.NewClient(addr)

	// constrcut transaction request
	req := api.CreateTransactionRequest{
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.1",
		},
		ChangeAddress:     nil,
		IgnoreUnconfirmed: false,
		Wallet: api.CreateTransactionRequestWallet{
			ID: wn,
		},
		To: targets,
	}

	response, err := client.CreateTransaction(req)
	if err != nil {
		panic(err)
	}

	itr, err := client.InjectTransaction(response.EncodedTransaction)
	if err != nil {
		panic(err)
	}

	fmt.Println(itr)

	return nil
}
