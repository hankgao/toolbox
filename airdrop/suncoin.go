package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/skycoin/skycoin/src/visor"
)

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

func getUxPool() map[string]visor.ReadableTransactionOutput {
	uxpool := make(map[string]visor.ReadableTransactionOutput)

	for i := 1; i <= maximumSeq; i++ {
		block := getBlock(i)

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

func getBlock(which int) visor.ReadableBlock {

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/block?seq=%d", webInterfacePort, which))
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
		fmt.Printf("%s", line)
		if err != nil {
			break
		}

		a := strings.Split(line, ",")
		balance, err := strconv.ParseFloat(a[2], 64)

		book[count] = addrItem{
			a[1],
			balance,
		}

		count++
	}

	return book, count
}
