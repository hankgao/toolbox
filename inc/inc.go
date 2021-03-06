// inc - Issue New Cryptocurrency
// Usage:
// inc <coin name>
// coin name is an optional parameter, when provided, the program will assume that the configuration.md file
// has been created, which will be parsed to create skycoin.go.sed and electron-main.js.sed
// In most cases, this is to upgrade to the newer version

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
	c := CoinConfigT{}
	if len(os.Args) > 1 {
		c = loadConfiguration(os.Args[1])
	} else {
		c = createConfiguration()
	}

	fmt.Println("Creating skycoin.go.sed")
	skycoinSed(filepath.Join("newcoins", strings.ToLower(c.CoinName), "skycoin.go.sed"), c)

	fmt.Println("Creating electron-main.js.sed")
	eletronMainJs(filepath.Join("newcoins", strings.ToLower(c.CoinName), "electron-main.js.sed"), c)
}

func createConfiguration() CoinConfigT {
	c := CoinConfigT{}

	r := bufio.NewReader(os.Stdin)

	fmt.Print("Coin name: ")
	c.CoinName, _ = r.ReadString('\n')
	c.CoinName = strings.TrimSuffix(c.CoinName, "\n")

	fmt.Print("Coin symbol: ")
	c.CoinSymbol, _ = r.ReadString('\n')
	c.CoinSymbol = strings.TrimSuffix(c.CoinSymbol, "\n")

	fmt.Print("Genesis coin volume (in droplets) : ")
	c.GenesisCoinVolume, _ = r.ReadString('\n')
	c.GenesisCoinVolume = strings.TrimSuffix(c.GenesisCoinVolume, "\n")
	c.GenesisCoinVolume = convertSn(c.GenesisCoinVolume)

	fmt.Print("Port: ")
	c.Port, _ = r.ReadString('\n')
	c.Port = strings.TrimSuffix(c.Port, "\n")

	fmt.Print("Web interface port: ")
	c.WebInterfacePort, _ = r.ReadString('\n')
	c.WebInterfacePort = strings.TrimSuffix(c.WebInterfacePort, "\n")

	// fmt.Print("RPC interface port: ")
	// c.RPCInterfacePort, _ = r.ReadString('\n')
	// c.RPCInterfacePort = strings.TrimSuffix(c.RPCInterfacePort, "\n")
	c.RPCInterfacePort = "XXXX"

	fmt.Print("Node 1 IP: ")
	c.Nodes[0], _ = r.ReadString('\n')
	c.Nodes[0] = strings.TrimSuffix(c.Nodes[0], "\n")

	fmt.Print("Node 2 IP: ")
	c.Nodes[1], _ = r.ReadString('\n')
	c.Nodes[1] = strings.TrimSuffix(c.Nodes[1], "\n")

	fmt.Print("Output folder: ")
	outputFolder, _ := r.ReadString('\n')
	outputFolder = strings.TrimSuffix(outputFolder, "\n")

	outputFolder = filepath.Join(outputFolder, c.CoinName)
	err := os.MkdirAll(outputFolder, 0755)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generating configurations ...")
	fillupCoinConfig(&c)

	fmt.Println("Creating configurations.md ...")
	configurationFile(filepath.Join(outputFolder, "configurations.md"), c)

	return c
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
	text = strings.Replace(text, "$TimestampReadable", c.GenesisTimestampReadable, -1)
	text = strings.Replace(text, "$GenesisCoinVolume", c.GenesisCoinVolume, -1)
	text = strings.Replace(text, "$PeerListURL", c.PeerListURL, -1)
	text = strings.Replace(text, "$Port", c.Port, -1)
	text = strings.Replace(text, "$WebInterfacePort", c.WebInterfacePort, -1)
	text = strings.Replace(text, "$RPCInterfacePort", c.RPCInterfacePort, -1)
	text = strings.Replace(text, "$GenesisUxID", c.GenesisUxID, -1)

	for i := 0; i < 2; i++ {
		text = strings.Replace(text, fmt.Sprintf("$node%03d", i+1), c.Nodes[i], -1)
		oldText := fmt.Sprintf("$nodewithport%03d", i+1)
		newText := fmt.Sprintf("%s:%s", c.Nodes[i], c.Port)
		text = strings.Replace(text, oldText, newText, -1)
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

// convertSn converts scientific notation to normal
func convertSn(sn string) string {
	a := strings.Split(sn, "e")
	if len(a) == 2 {
		firstPart := a[0]
		zeros, err := strconv.Atoi(a[1])
		if err != nil {
			return sn
		}

		return firstPart + strings.Repeat("0", zeros)
	}

	return sn

}

// loadConfiguration
func loadConfiguration(cn string) CoinConfigT {
	config := CoinConfigT{}

	fn := filepath.Join("newcoins/", cn, "/configurations.md")
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	kvMap := make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '|' {
			tokens := strings.Split(line[1:], "|")
			if len(tokens) >= 2 {
				key := strings.TrimSpace(tokens[0])

				if key == "" || key == "Item" || key[0] == '-' {
					continue
				}

				value := strings.TrimSpace(tokens[1])

				kvMap[key] = value
			}
		}
	}

	config.BlockchainPubkeyStr = kvMap["BlockchainPubkeyStr"]
	config.BlockchainSeckeyStr = kvMap["BlockchainSeckeyStr"]
	config.CoinName = kvMap["Coin name"]
	config.CoinSymbol = kvMap["Coin symbol"]
	config.gasPrivateKey = kvMap["PrivateKey"]
	config.GenesisAddressStr = kvMap["GenesisAddressStr"]
	config.GenesisCoinVolume = kvMap["GenesisCoinVolume"]
	config.GenesisSignatureStr = kvMap["GenesisSignatureStr"]

	config.GenesisTimestamp, err = strconv.ParseInt(kvMap["GenesisTimestamp"], 10, 64)
	config.GenesisUxID = kvMap["GenesisUxID"]
	config.Nodes[0] = kvMap["Node 1"]
	config.Nodes[1] = kvMap["Node 2"]
	config.Port = kvMap["Port"]
	config.RPCInterfacePort = kvMap["RPCInterfacePort"]
	config.WebInterfacePort = kvMap["WebInterfacePort"]

	return config
}
