// Package snapshot prints out the whole book, by address
//
// Usage:
//   snapshot <coin name> <max block heigth>
// Example:
//   snapshot shellcoin 100
//
// Note that you need to run a wallet or a node to ensure this gadget works
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hankgao/toolbox/src/skytb"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf(`Correct usage: 
			snapshot <coin name> <max block height> 
			for example: snapshot skycoin 10000
			`)
		return
	}

	coinName := os.Args[1]
	maxBlock, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("Max block height has to be an integer\n")
	}

	skytb.Snapshot(coinName, maxBlock)
}
