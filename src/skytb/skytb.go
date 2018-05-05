package skytb

import (
	"fmt"
	"log"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	bip39 "github.com/tyler-smith/go-bip39"
)

// GenerateSeed generates a seed (12 random words)
func GenerateSeed() string {
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

// GetFirstAddress returns the first address created from a seed
func GetFirstAddress(seed string) (string, string) {
	publicK, privateK := cipher.GenerateDeterministicKeyPair([]byte(seed))
	return cipher.AddressFromPubKey(publicK).String(), privateK.Hex()

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
		tx.PushOutput(addr, each*1e6, 1)
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
