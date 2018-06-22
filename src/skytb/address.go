package skytb

import (
	"github.com/skycoin/skycoin/src/cipher"
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
