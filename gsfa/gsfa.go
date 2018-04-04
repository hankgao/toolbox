// gsfa = Genrate seed and the first address created from this seed
// you can choose how many of them are created
// Sample
// $ go run gsfa -qty=100
package main

import (
	"flag"
	"fmt"

	"github.com/skycoin/skycoin/src/cipher"
	bip39 "github.com/tyler-smith/go-bip39"
)

func main() {
	var qty int
	flag.IntVar(&qty, "qty", 1, "number of seeds & addresses to generate")
	flag.Parse()

	for i := 0; i < qty; i++ {
		seed := generateSeed()
		addr := createFirstAddr(seed)
		fmt.Printf("%s, %s\n", seed, addr)
	}

}

func generateSeed() string {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}

	seed, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}

	return seed
}

func createFirstAddr(seed string) string {
	publicK, _ := cipher.GenerateDeterministicKeyPair([]byte(seed))
	return cipher.AddressFromPubKey(publicK).String()

}
