// inc - Issue New Cryptocurrency

package main

import (
	"bufio"
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
	fmt.Println("Creating ")
	c := CoinConfigT{}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Coin name: ")
	c.CoinName, _ = r.ReadString('\n')
	c.CoinName = strings.TrimSuffix(c.CoinName, "\n")

	fmt.Print("Coin symbol: ")
	c.CoinSymbol, _ = r.ReadString('\n')
	c.CoinSymbol = strings.TrimSuffix(c.CoinSymbol, "\n")

	fmt.Print("Genesis coin volume : ")
	c.GenesisCoinVolume, _ = r.ReadString('\n')
	c.GenesisCoinVolume = strings.TrimSuffix(c.GenesisCoinVolume, "\n")

	fmt.Print("Port: ")
	c.Port, _ = r.ReadString('\n')
	c.Port = strings.TrimSuffix(c.Port, "\n")

	fmt.Print("Web interface port: ")
	c.WebInterfacePort, _ = r.ReadString('\n')
	c.WebInterfacePort = strings.TrimSuffix(c.WebInterfacePort, "\n")

	fmt.Print("RPC interface port: ")
	c.RPCInterfacePort, _ = r.ReadString('\n')
	c.RPCInterfacePort = strings.TrimSuffix(c.RPCInterfacePort, "\n")

	fmt.Print("Node 1 IP: ")
	c.Nodes[0], _ = r.ReadString('\n')
	c.Nodes[0] = strings.TrimSuffix(c.Nodes[0], "\n")

	fmt.Print("Node 2 IP: ")
	c.Nodes[1], _ = r.ReadString('\n')
	c.Nodes[1] = strings.TrimSuffix(c.Nodes[1], "\n")

	fmt.Println("Generating configurations ...")
	fillupCoinConfig(&c)

	fmt.Println("Creating configuration.md ...")
	configurationFile("configuration.md", c)

	fmt.Println("Creating skycoin.go.sed")
	skycoinSed("skycoin.go.sed", c)
}

func fillupCoinConfig(c *CoinConfigT) {

	publicK, privateK := cipher.GenerateKeyPair()

	c.BlockchainPubkeyStr = publicK.Hex()
	c.BlockchainSeckeyStr = privateK.Hex()

	c.gasSeed = skytb.GenerateSeed()
	c.GenesisAddressStr, c.gasPrivateKey = skytb.GetFirstAddress(c.gasSeed)

	now := time.Now()
	c.GenesisTimestamp = now.Unix()

	now.Format(time.RFC822)
	c.GenesisTimestampReadable = now.String()

	c.PeerListURL = ""
	c.GenesisUxID = ""

	// List 100 addresseed used for initial distribution
	// TODO: write to a csv file
	addrs := make([]string, 100)
	for i := 0; i < 100; i++ {
		seed := skytb.GenerateSeed()
		address, _ := skytb.GetFirstAddress(seed)
		c.DistributionAddresses[i] = fmt.Sprintf("- %03d, %40s, %s\n", i+1, address, seed)

		addrs = append(addrs, address)
	}

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
		oldText := fmt.Sprintf("$nodewithport%03d", i+1)
		newText := fmt.Sprintf("%s:%s", c.Nodes[i], c.WebInterfacePort)
		text = strings.Replace(text, oldText, newText, -1)
	}

	fmt.Println(text)

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
