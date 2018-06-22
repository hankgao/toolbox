package skytb

import (
	"fmt"
	"log"

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

// SpendFromWallet spends coins from a specified wallet
func SpendFromWallet(wlt string, dst []Balance) error {
	return nil
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
