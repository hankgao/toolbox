package main

// CoinPair represents a tradiing pair
type CoinPair struct {
	FromCoin    string  `json:"from_coin"`
	ToCoin      string  `json:"to_coin"`
	ExhangeRate float64 `json:"exchange_rate"`
	ValidFrom   string
	ValidUntil  string
}

// CoinPairsR contains all coin pairs
type CoinPairsR struct {
	Number    int        `json:"number"`
	CoinPairs []CoinPair `jason:"coin_pairs"`
}

// Order reprensents an order
type Order struct {
	CoinFrom        string `json:"coin_from"`
	CoinTo          string `json:"coin_to"`
	ReceivingAddr   string `json:"receiving_address"`
	DepositingAddr  string `json:"despoit_address"` // is this property necessary?
	InvestorName    string `json:"investor_name"`
	InvestorMobile  string `json:"investor_mobile"`
	InvestorAddress string `json:"investor_address"`
	TimePlaced      string `json:"time_placed"`
}
