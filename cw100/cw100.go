// cw100 creates 100 wallets from 100 seeds
// usage:
//   cw100 <seeds_file_name>
// seed file has to be in the following format:
// give nurse mechanic lock swamp guard blue young bike tennis reject orchard
// cruel diesel broccoli fault argue broom cabin veteran pact local pet absurd

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hankgao/toolbox/src/skytb"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf(`
			Usage: 
				cw100 <coin_name> <seed_file_name> 
				`)
		return
	}

	cn := os.Args[1]
	sfn := os.Args[2]

	seeds := read100Seeds(sfn)

	skytb.CreateDistributionWallets(cn, seeds)
}

// read100Seeds reads 100 seeds from a file
func read100Seeds(fn string) []string {
	seeds := make([]string, 100)

	f, err := os.Open(fn)
	defer f.Close()

	reader := bufio.NewReader(f)

	if err != nil {
		panic(err)
	}

	seed := ""
	for index := 0; ; index++ {
		seed, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		seeds[index] = seed[0 : len(seed)-1] // remove the trailing carraige
	}

	return seeds
}
