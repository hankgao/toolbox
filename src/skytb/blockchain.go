package skytb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/skycoin/skycoin/src/visor"
)

// contians code that analyze a blockchain

// Snapshot prints snapshot of a blockchain to stdout, in CSV format
func Snapshot(coinName string, maximumSeq int) {
	uxpool := GetUxPool(coinName, maximumSeq)
	book := aggregateAddresses(uxpool)
	writeBook(book)
}

// writeBook prints book to stdout in CSV format, which can be redirected to any file
func writeBook(book map[string]float64) {
	for addr, coins := range book {
		fmt.Printf("%s,%f\n", addr, coins)
	}
}

// aggregateAddresses combine all transaction outputs by addresses
func aggregateAddresses(uxpool map[string]visor.ReadableTransactionOutput) map[string]float64 {
	addressBook := make(map[string]float64)

	for _, ux := range uxpool {
		coins, err := strconv.ParseFloat(ux.Coins, 64)

		if err != nil {
			panic(err)
		}

		addressBook[ux.Address] += coins
	}

	return addressBook
}

// GetUxPool create UNSPENT transaction output pool
func GetUxPool(coinName string, maximumSeq int) map[string]visor.ReadableTransactionOutput {
	uxpool := make(map[string]visor.ReadableTransactionOutput)

	for i := 1; i <= maximumSeq; i++ {
		block := getBlock(coinName, i)

		for _, t := range block.Body.Transactions {
			// remove input
			for _, input := range t.In {
				delete(uxpool, input)
			}

			// Record outputs
			for _, output := range t.Out {
				uxpool[output.Hash] = output
			}
		}
	}

	return uxpool
}

// getBlock returns block at heigth of which
func getBlock(coinName string, which int) visor.ReadableBlock {

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/block?seq=%d", CoinTypesSupported[coinName].WebInterfacePort, which))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyStr, err := ioutil.ReadAll(resp.Body)

	block := visor.ReadableBlock{}

	err = json.Unmarshal(bodyStr, &block)

	if err != nil {
		panic(err)
	}

	return block

}
