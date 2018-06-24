package skytb

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

}
