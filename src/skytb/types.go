package skytb

import "strings"

// Balance represents balance
type Balance struct {
	addr  string
	coins uint64 // in droplest
	hours uint64
}

// CoinType represents a skyledger based coin
type CoinType struct {
	CoinName         string // all lowercase, like shellcoin, mzcoin etc
	CoinSymbol       string
	Port             int
	WebInterfacePort int
	RPCInterfacePort int
}

// AddrItem represents a pair of address and balance
type AddrItem struct {
	Addr    string
	Balance uint64 // in droplet
}

// CoinTypesSupported stores all cointypes that have been implemented so far
var CoinTypesSupported map[string]CoinType

func init() {
	CoinTypesSupported = make(map[string]CoinType)
	CoinTypesSupported["mzcoin"] = CoinType{
		CoinName:         "mzcoin",
		CoinSymbol:       "MZC",
		Port:             7000,
		WebInterfacePort: 7420,
		RPCInterfacePort: 7430,
	}

	CoinTypesSupported["shellcoin"] = CoinType{
		CoinName:         "shellcoin",
		CoinSymbol:       "SC2",
		Port:             7100,
		WebInterfacePort: 7520,
		RPCInterfacePort: 7530,
	}

	CoinTypesSupported["metalicoin"] = CoinType{
		CoinName:         "metalicoin",
		CoinSymbol:       "MTC",
		Port:             7400,
		WebInterfacePort: 7820,
		RPCInterfacePort: 7530,
	}

	CoinTypesSupported["metalcoin"] = CoinType{
		CoinName:         "metalcoin",
		CoinSymbol:       "MTC",
		Port:             7240,
		WebInterfacePort: 7250,
		RPCInterfacePort: 7260,
	}

	CoinTypesSupported["angelcoin"] = CoinType{
		CoinName:         "angelcoin",
		CoinSymbol:       "AGLC",
		Port:             7480,
		WebInterfacePort: 7490,
		RPCInterfacePort: 0,
	}
}

// IsCoinTypeSupported check whether a coin type is supported
func IsCoinTypeSupported(coinType string) bool {
	_, ok := CoinTypesSupported[coinType]
	return ok
}

// CoinTypeDetails return CoinType structure, given its coin type name
func CoinTypeDetails(coinType string) CoinType {
	ct := CoinType{}
	if IsCoinTypeSupported(coinType) {
		ct = CoinTypesSupported[coinType]
	}

	return ct
}

// RemoveGarbage removes garbage letters from a string
func RemoveGarbage(s string) string {
	// rempve space, '\n', ','
	return strings.Map(func(r rune) rune {
		if r == 32 || r == '\n' || r == ',' || r == '\r' {
			return -1
		}
		return r
	}, s)
}
