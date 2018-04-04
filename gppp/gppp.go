// Package main generate a public-private key pair
package main

import (
	"fmt"

	"github.com/skycoin/skycoin/src/cipher"
)

func main() {

	publicK, privateK := cipher.GenerateKeyPair()

	publicKey := publicK.Hex()
	privateKey := privateK.Hex()

	fmt.Printf("Public  key: %s\n", publicKey)
	fmt.Printf("Private key: %s\n", privateKey)

}
