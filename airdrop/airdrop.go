// airdrop distributes coins to addresses listed in a book file in CSV format.
// In most cases, this is to air drop coins to users who already have coins of another type
// airdrop has two parameters
//    - coin names: coin type that will be dropped
//    - book: a CSV file containing addresses with amounts of coins to be dropped
//
// note:
//   you need to run cw100 to create distribution wallets before you can actually distribute coins
//
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hankgao/toolbox/src/skytb"
	"github.com/skycoin/skycoin/src/api"
)

func readBook(fn string) ([]api.Receiver, int) {

	book := []api.Receiver{}

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

		line = strings.TrimSuffix(line, "\r\n")

		a := strings.Split(line, ",")
		addr := a[0]
		coins := a[1]

		if err != nil {
			panic(err)
		}

		book = append(book, api.Receiver{
			Address: addr,
			Coins:   coins,
		})

		count++

	}

	return book, count
}

func main() {

	if len(os.Args) < 3 {
		fmt.Printf(`
			Usage: 
				airdrop <metalicoin.1414.csv> <metalcoin>
				`)
		return
	}

	targets, _ := readBook(os.Args[2] /* book */)
	skytb.DistributeCoins(os.Args[1] /* coin name */, "airdrop.wlt", targets)
}
