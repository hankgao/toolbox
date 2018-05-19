// inc - Issue New Cryptocurrency

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gitee.com/hankgaosh/toolbox/src/skytb"
	"github.com/skycoin/skycoin/src/cipher"
)

const version = "0.1.1"

// CoinConfigT contains configurations for a new coin
type CoinConfigT struct {
	CoinName                 string
	CoinSymbol               string
	BlockchainPubkeyStr      string
	BlockchainSeckeyStr      string
	GenesisAddressStr        string
	gasPrivateKey            string
	gasSeed                  string
	GenesisSignatureStr      string
	GenesisTimestamp         int64
	GenesisTimestampReadable string
	GenesisCoinVolume        string
	PeerListURL              string
	Port                     string
	WebInterfacePort         string
	RPCInterfacePort         string
	GenesisUxID              string
	Nodes                    [4]string
	DistributionAddresses    [100]string // in format: 001,seed,address
}

func main() {
	c := newCoinConfig()
	c.CoinName = "LiquorCoin"
	c.CoinSymbol = "LQC"

	c.Nodes[0] = "127.324.234.234"
	c.Nodes[1] = "234.324.234.23"

	c.Port = "8000"
	c.WebInterfacePort = "8420"
	c.RPCInterfacePort = "8430"

	configurationFile("configuration.md", c)
	skycoinSed("skycoin.go.sed", c)
}

func newCoinConfig() CoinConfigT {
	c := CoinConfigT{}

	publicK, privateK := cipher.GenerateKeyPair()

	c.BlockchainPubkeyStr = publicK.Hex()
	c.BlockchainSeckeyStr = privateK.Hex()

	c.gasSeed = skytb.GenerateSeed()
	c.GenesisAddressStr, c.gasPrivateKey = skytb.GetFirstAddress(c.gasSeed)

	now := time.Now()
	c.GenesisTimestamp = now.Unix()

	now.Format(time.RFC822)
	c.GenesisTimestampReadable = now.String()

	c.GenesisCoinVolume = "300e12"
	c.PeerListURL = ""
	c.Port = ""
	c.WebInterfacePort = ""
	c.RPCInterfacePort = ""
	c.GenesisUxID = ""
	c.Nodes[0] = ""
	c.Nodes[1] = ""
	c.Nodes[2] = ""
	c.Nodes[3] = ""

	// List 100 addresseed used for initial distribution
	// TODO: write to a csv file
	addrs := make([]string, 100)
	for i := 0; i < 100; i++ {
		seed := skytb.GenerateSeed()
		address, _ := skytb.GetFirstAddress(seed)
		c.DistributionAddresses[i] = fmt.Sprintf("- %03d, %40s, %s\n", i+1, address, seed)

		addrs = append(addrs, address)
	}

	return c
}

func injectValues(text string, c CoinConfigT) string {
	text = strings.Replace(text, "$CoinName", c.CoinName, -1)
	text = strings.Replace(text, "$coinname", strings.ToLower(c.CoinName), -1)
	text = strings.Replace(text, "$Coinname", strings.Title(strings.ToLower(c.CoinName)), -1)
	text = strings.Replace(text, "$CoinSymbol", c.CoinSymbol, -1)
	text = strings.Replace(text, "$BlockchainPubkeyStr", c.BlockchainPubkeyStr, -1)
	text = strings.Replace(text, "$BlockchainSeckeyStr", c.BlockchainSeckeyStr, -1)
	text = strings.Replace(text, "$GenesisAddressStr", c.GenesisAddressStr, -1)
	text = strings.Replace(text, "$gasPrivateKey", c.gasPrivateKey, -1)
	text = strings.Replace(text, "$gasSeed", c.gasSeed, -1)
	text = strings.Replace(text, "$GenesisSignatureStr", c.GenesisSignatureStr, -1)
	text = strings.Replace(text, "$GenesisTimestamp", fmt.Sprintf("%d", c.GenesisTimestamp), -1)
	text = strings.Replace(text, "$GenesisTimestampReadable", c.GenesisTimestampReadable, -1)
	text = strings.Replace(text, "$GenesisCoinVolume", c.GenesisCoinVolume, -1)
	text = strings.Replace(text, "$PeerListURL", c.PeerListURL, -1)
	text = strings.Replace(text, "$Port", c.Port, -1)
	text = strings.Replace(text, "$WebInterfacePort", c.WebInterfacePort, -1)
	text = strings.Replace(text, "$RPCInterfacePort", c.RPCInterfacePort, -1)
	text = strings.Replace(text, "$GenesisUxID", c.GenesisUxID, -1)

	for i := 0; i < 2; i++ {
		text = strings.Replace(text, fmt.Sprintf("$node%03d", i+1), c.Nodes[i], -1)
		text = strings.Replace(text, fmt.Sprintf("$node%03d:port", i+1),
			fmt.Sprintf("%s:%s", c.Nodes[i], c.WebInterfacePort), -1)
	}

	return text

}

func write2File(fn, text string) {
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(text + "\n")
}
