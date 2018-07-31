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
	"strconv"
	"strings"

	"github.com/hankgao/toolbox/src/skytb"
)

func readBook(fn string) ([]skytb.AddrItem, int) {
	book := make([]skytb.AddrItem, 2000)
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

		bs := skytb.RemoveGarbage(a[1])

		balancef, _ := strconv.ParseFloat(bs, 64)
		balance := uint64(balancef * 1e6)

		// last three digits have to be 000!
		if balance%1000 > 0 {
			balance = balance/1000*1000 + 1000
		}

		if err != nil {
			panic(err)
		}

		book[count] = skytb.AddrItem{
			Addr:    a[0],
			Balance: balance,
		}

		// if strings.Contains(book[count].addr, "4K") {
		// 	fmt.Printf("%s=>%d\n", book[count].addr, book[count].balance)

		// }

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
	book := os.Args[1]
	coinName := os.Args[2]

	accouts, count := readBook(book)
	skytb.DistCoinsNamedWallets(coinName, accouts, count)

}
