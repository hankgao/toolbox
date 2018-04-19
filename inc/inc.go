// inc - Issue New Cryptocurrency

package main

import (
	"fmt"
	"time"

	"gitee.com/hankgaosh/toolbox/src/skytb"
	"github.com/skycoin/skycoin/src/cipher"
)

func main() {
	publicK, privateK := cipher.GenerateKeyPair()

	publBlockchainPubkeyStricKey := publicK.Hex()
	BlockchainSeckeyStr := privateK.Hex()

	fmt.Printf("publBlockchainPubkeyStr :%s\n", publBlockchainPubkeyStricKey)
	fmt.Printf("BlockchainSeckeyStr :%s\n", BlockchainSeckeyStr)

	seed := skytb.GenerateSeed()
	GenesisAddressStr := skytb.GetFirstAddress(seed)
	fmt.Printf("GenesisAddressStr :%s(%s)\n", GenesisAddressStr, seed)

	now := time.Now()
	GenesisTimestamp := now.Unix()
	fmt.Printf("GenesisTimestamp %d(%s)\n", GenesisTimestamp, now.String())

	// List 100 addresseed used for initial distribution
	// TODO: write to a csv file
	addrs := make([]string, 100)
	for i := 0; i < 100; i++ {
		seed := skytb.GenerateSeed()
		address := skytb.GetFirstAddress(seed)
		fmt.Printf("%03d,%s,%s\n", i, address, seed)

		addrs = append(addrs, address)
	}

	// Get signature for the initial distribution to 100 addresses

	// GenesisSignatureStr
	// GenesisCoinVolume
	// PeerListURL
	// WebInterfacePort
	// RPCInterfacePort
	// Port
}
